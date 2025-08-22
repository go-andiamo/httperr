package httperr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"slices"
	"strings"
)

// StatusError is an error interface that supports status codes
type StatusError interface {
	error
	// StatusCode returns the HTTP status code for the error
	StatusCode() int
}

type HttpError interface {
	error
	StatusError
	// WithCause returns a HttpError with the cause set
	WithCause(cause error) HttpError
	Unwrap() error
	Cause() error
	// StackInfo returns the call stack info for the error
	StackInfo() StackInfo
	// AddHeaders adds the supplied response headers to the error
	AddHeaders(hdrs map[string]string) HttpError
	// AddHeader adds the supplied response header to the error
	AddHeader(header string, value string) HttpError
	// Headers returns the additional response headers for the HttpError
	//
	// additional headers are written to the response by DefaultErrorWriter
	Headers() map[string]string
	// AddReasons adds the supplied reasons to the error
	AddReasons(reasons ...any) HttpError
	// AddReason adds the supplied reason to the error
	AddReason(reasons any) HttpError
	// Reasons returns the reasons for the error
	//
	// reasons are written to the response body by the DefaultErrorWriter
	Reasons() []any
	// Write writes the error to the http.ResponseWriter
	//
	// it uses the DefaultErrorWriter - if DefaultErrorWriter is nil, just the status
	// code and any additional headers are written to the response writer
	Write(w http.ResponseWriter)
}

// New creates a new HttpError for the specified status code with stack info
//
// if the msg arg is an empty string, the message is derived from http.StatusText for the status code
func New(status int, msg string) HttpError {
	return newError(status, msg, nil, getStackInfo())
}

// Newf creates a new HttpError for the specified status code with stack info and a formatted message
//
// if the formatted message is an empty string, the message is derived from http.StatusText for the status code
func Newf(status int, format string, a ...any) HttpError {
	return newError(status, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// Wrap wraps an existing error with a HttpError
//
// if the DefaultErrorStatusResolver is set, the status will be determined using that resolver
//
// Note: the stack info is based on the point at which Wrap is called (rather than the callers of the wrapped error)
func Wrap(cause error, defaultStatus int) HttpError {
	if cause == nil {
		return nil
	}
	if DefaultErrorStatusResolver == nil {
		return newError(defaultStatus, "", cause, getStackInfo())
	}
	return newError(DefaultErrorStatusResolver.Resolve(cause, defaultStatus), "", cause, getStackInfo())
}

func newError(status int, msg string, cause error, si StackInfo) HttpError {
	if msg == "" {
		msg = http.StatusText(status)
	}
	return &httpError{
		message: msg,
		stack:   si,
		cause:   cause,
		status:  status,
		headers: make(map[string]string),
	}
}

type httpError struct {
	message string
	stack   StackInfo
	cause   error
	status  int
	reasons []any
	headers map[string]string
}

var _ error = (*httpError)(nil)
var _ HttpError = (*httpError)(nil)
var _ fmt.Formatter = (*httpError)(nil)

func (e *httpError) MarshalJSON() ([]byte, error) {
	m := map[string]any{
		ptyError: e.message,
	}
	if e.reasons != nil {
		m[ptyReasons] = e.reasons
	}
	if e.cause != nil && DefaultErrorWriterShowCause {
		m[ptyCause] = e.cause.Error()
	}
	if len(e.stack) > 0 && DefaultErrorWriterShowStack {
		info := make([]string, len(e.stack))
		for i, f := range e.stack {
			info[i] = fmt.Sprintf("%s:%d", f.Function, f.Line)
		}
		m[ptyStack] = info
	}
	return json.Marshal(m)
}

func (e *httpError) StatusCode() int {
	return e.status
}

func (e *httpError) Write(w http.ResponseWriter) {
	if DefaultErrorWriter == nil {
		for k, v := range e.headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(e.status)
		return
	}
	DefaultErrorWriter.WriteError(e, w)
}

func (e *httpError) Error() string {
	return e.message
}

func (e *httpError) Unwrap() error {
	return e.cause
}

func (e *httpError) Cause() error {
	return e.cause
}

func (e *httpError) WithCause(cause error) HttpError {
	e.cause = cause
	return e
}

func (e *httpError) Reasons() []any {
	return e.reasons
}

func (e *httpError) AddReasons(reasons ...any) HttpError {
	e.reasons = append(e.reasons, reasons...)
	return e
}

func (e *httpError) AddReason(reason any) HttpError {
	e.reasons = append(e.reasons, reason)
	return e
}

func (e *httpError) AddHeaders(hdrs map[string]string) HttpError {
	for k, v := range hdrs {
		e.headers[k] = v
	}
	return e
}

func (e *httpError) AddHeader(header string, value string) HttpError {
	e.headers[header] = value
	return e
}

func (e *httpError) Headers() map[string]string {
	return e.headers
}

func (e *httpError) StackInfo() StackInfo {
	return e.stack
}

func (e *httpError) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		_, _ = fmt.Fprintf(f, "%s", e.message)
		if f.Flag('+') {
			if e.cause != nil {
				_, _ = fmt.Fprintf(f, ": %+v", e.cause)
			}
			if len(e.stack) > 0 && DefaultFrameFormatter != nil {
				_, _ = io.WriteString(f, DefaultFrameFormatter.StartLine())
				for _, fr := range e.stack {
					_, _ = io.WriteString(f, DefaultFrameFormatter.FrameLine(fr))
				}
			}
		} else if e.cause != nil {
			_, _ = fmt.Fprintf(f, ": %v", e.cause)
		}
	case 's':
		_, _ = io.WriteString(f, e.message)
	case 'q':
		_, _ = fmt.Fprintf(f, "%q", e.message)
	default:
		_, _ = io.WriteString(f, "%!")
		_, _ = io.WriteString(f, string(verb))
		_, _ = io.WriteString(f, "(httperr.HttpError)")
	}
}

type StackInfo []runtime.Frame

func getStackInfo() StackInfo {
	result := make(StackInfo, 0, MaxStackDepth)
	const skip = 3
	pc := make([]uintptr, MaxStackDepth)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	for frame, more := frames.Next(); more && len(result) < int(MaxStackDepth); frame, more = frames.Next() {
		if DefaultPackageFilter != nil || DefaultPackageName != "" {
			if !packageMatch(packageFromFunction(frame.Function)) {
				continue
			}
		}
		result = append(result, frame)
	}
	return result
}

func packageMatch(full string, short string, parts []string) bool {
	result := true
	if DefaultPackageFilter != nil && !DefaultPackageFilter.Include(full) {
		result = false
	}
	if result && DefaultPackageName != "" {
		if strings.HasSuffix(DefaultPackageName, "/") {
			result = slices.Contains(parts, DefaultPackageName[:len(DefaultPackageName)-1])
		} else {
			result = DefaultPackageName == short
		}
	}
	return result
}

func packageFromFunction(name string) (full string, short string, parts []string) {
	full = name
	if s := strings.LastIndexByte(full, '/'); s >= 0 {
		if d := strings.IndexByte(name[s+1:], '.'); d >= 0 {
			short = name[s+1 : s+d+1]
		}
		full = full[:s+1] + short
	} else if d := strings.IndexByte(name, '.'); d >= 0 {
		full = full[:d]
		short = full
	}
	return full, short, strings.Split(full, "/")
}
