package httperr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// DefaultErrorWriter is the default error writer used by HttpError.Write
var DefaultErrorWriter ErrorWriter = &errorWriter{}

// DefaultErrorWriterShowCause determines whether the internal default error writer
// shows the cause in the response body json
var DefaultErrorWriterShowCause = false

// DefaultErrorWriterShowStack determines whether the internal default error writer
// shows the stack trace in the response body json
var DefaultErrorWriterShowStack = false

// ErrorWriter is the interface used to write errors (i.e. HttpError.Write)
type ErrorWriter interface {
	WriteError(e error, w http.ResponseWriter)
}

type errorWriter struct{}

var _ ErrorWriter = (*errorWriter)(nil)

const (
	ptyError        = "$error"
	ptyReasons      = "$reasons"
	ptyCause        = "$cause"
	prtStack        = "$stack"
	hdrContentType  = "Content-Type"
	applicationJson = "application/json"
)

func (ew *errorWriter) WriteError(err error, w http.ResponseWriter) {
	w.Header().Set(hdrContentType, applicationJson)
	status := http.StatusInternalServerError
	body := map[string]any{
		ptyError: http.StatusText(http.StatusInternalServerError),
	}
	switch et := err.(type) {
	case HttpError:
		status = et.StatusCode()
		body[ptyError] = err.Error()
		if DefaultErrorWriterShowStack {
			if stack := et.StackInfo(); len(stack) > 0 {
				info := make([]string, len(stack))
				for i, f := range stack {
					info[i] = fmt.Sprintf("%s:%d", f.Function, f.Line)
				}
				body[prtStack] = info
			}
		}
		if reasons := et.Reasons(); len(reasons) > 0 {
			body[ptyReasons] = reasons
		}
		for k, v := range et.Headers() {
			w.Header().Set(k, v)
		}
	case StatusError:
		status = et.StatusCode()
		if err.Error() != "" {
			body[ptyError] = err.Error()
		} else {
			body[ptyError] = http.StatusText(status)
		}
	case error:
		body[ptyError] = et.Error()
	}
	if DefaultErrorWriterShowCause {
		if cause := errors.Unwrap(err); cause != nil {
			body[ptyCause] = cause.Error()
		}
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
