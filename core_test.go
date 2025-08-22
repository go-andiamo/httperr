package httperr

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	e := New(http.StatusBadRequest, "")
	require.Error(t, e)
	require.Equal(t, "Bad Request", e.Error())
	require.Equal(t, http.StatusBadRequest, e.StatusCode())
	require.Empty(t, e.Reasons())
	require.Empty(t, e.Headers())
}

func TestNewf(t *testing.T) {
	e := Newf(http.StatusBadRequest, "something %d", 1)
	require.Error(t, e)
	require.Equal(t, "something 1", e.Error())
	require.Equal(t, http.StatusBadRequest, e.StatusCode())

	e = Newf(http.StatusBadRequest, "")
	require.Error(t, e)
	require.Equal(t, "Bad Request", e.Error())
	require.Equal(t, http.StatusBadRequest, e.StatusCode())
}

func TestWrap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		e := Wrap(nil, http.StatusBadRequest)
		require.NoError(t, e)
	})
	t.Run("error", func(t *testing.T) {
		e := Wrap(errors.New("some error"), http.StatusInternalServerError)
		require.Error(t, e)
		require.Equal(t, http.StatusInternalServerError, e.StatusCode())
		require.Equal(t, "Internal Server Error", e.Error())
		cause := errors.Unwrap(e)
		require.Error(t, cause)
		require.Equal(t, "some error", cause.Error())
	})
	t.Run("error (with resolver)", func(t *testing.T) {
		DefaultErrorStatusResolver = &testErrorStatusResolver{}
		defer func() { DefaultErrorStatusResolver = nil }()
		e := Wrap(sql.ErrNoRows, http.StatusInternalServerError)
		require.Error(t, e)
		require.Equal(t, http.StatusNotFound, e.StatusCode())
		require.Equal(t, "Not Found", e.Error())
		cause := errors.Unwrap(e)
		require.Error(t, cause)
		require.Equal(t, sql.ErrNoRows.Error(), cause.Error())
	})
	t.Run("error (with resolver, not resolved)", func(t *testing.T) {
		DefaultErrorStatusResolver = &testErrorStatusResolver{}
		defer func() { DefaultErrorStatusResolver = nil }()
		e := Wrap(errors.New("cause"), http.StatusInternalServerError)
		require.Error(t, e)
		require.Equal(t, http.StatusInternalServerError, e.StatusCode())
		require.Equal(t, "Internal Server Error", e.Error())
		cause := errors.Unwrap(e)
		require.Error(t, cause)
		require.Equal(t, "cause", cause.Error())
	})
}

type testErrorStatusResolver struct{}

var _ ErrorStatusResolver = (*testErrorStatusResolver)(nil)

func (t *testErrorStatusResolver) Resolve(err error, fallback int) int {
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound
	}
	return fallback
}

func TestError_Unwrap(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey")
	require.Error(t, e)
	e2 := e.Unwrap()
	require.NoError(t, e2)
	require.NoError(t, errors.Unwrap(e))
	e = e.WithCause(errors.New("cause"))
	e2 = e.Unwrap()
	require.Error(t, e2)
	require.Error(t, errors.Unwrap(e))
}

func TestError_Cause(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey")
	require.Error(t, e)
	e2 := e.Cause()
	require.NoError(t, e2)
	e = e.WithCause(errors.New("cause"))
	e2 = e.Cause()
	require.Error(t, e2)
}

func TestError_StackInfo(t *testing.T) {
	t.Run("with DefaultPackageName", func(t *testing.T) {
		DefaultPackageName = "httperr"
		defer func() {
			DefaultPackageName = ""
		}()
		ln := lineNumber() + 1
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		si := e.StackInfo()
		require.Len(t, si, 1)
		require.True(t, strings.HasPrefix(si[0].Function, "github.com/go-andiamo/httperr"))
		require.Contains(t, si[0].Function, "TestError_StackInfo")
		require.Equal(t, ln, si[0].Line)
	})
	t.Run("with DefaultPackageName (starts with)", func(t *testing.T) {
		DefaultPackageName = "httperr/"
		defer func() {
			DefaultPackageName = ""
		}()
		ln := lineNumber() + 1
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		si := e.StackInfo()
		require.Len(t, si, 1)
		require.True(t, strings.HasPrefix(si[0].Function, "github.com/go-andiamo/httperr"))
		require.Contains(t, si[0].Function, "TestError_StackInfo")
		require.Equal(t, ln, si[0].Line)
	})
	t.Run("with DefaultPackageFilter", func(t *testing.T) {
		DefaultPackageFilter = &testPackageFilter{}
		defer func() {
			DefaultPackageFilter = nil
		}()
		ln := lineNumber() + 1
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		si := e.StackInfo()
		require.Len(t, si, 1)
		require.True(t, strings.HasPrefix(si[0].Function, "github.com/go-andiamo/httperr"))
		require.Contains(t, si[0].Function, "TestError_StackInfo")
		require.Equal(t, ln, si[0].Line)
	})
	t.Run("with SetDefaultPackageFilter", func(t *testing.T) {
		SetDefaultPackageFilter("github.com/go-andiamo/httperr")
		defer func() {
			DefaultPackageFilter = nil
		}()
		ln := lineNumber() + 1
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		si := e.StackInfo()
		require.Len(t, si, 1)
		require.True(t, strings.HasPrefix(si[0].Function, "github.com/go-andiamo/httperr"))
		require.Contains(t, si[0].Function, "TestError_StackInfo")
		require.Equal(t, ln, si[0].Line)
	})
}

type testPackageFilter struct{}

var _ PackageFilter = (*testPackageFilter)(nil)

func (pf *testPackageFilter) Include(packageName string) bool {
	return strings.Contains(packageName, "httperr")
}

func TestError_Format(t *testing.T) {
	t.Run("v", func(t *testing.T) {
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		require.Equal(t, "fooey", fmt.Sprintf("%v", e))
	})
	t.Run("v with cause", func(t *testing.T) {
		e := New(http.StatusBadRequest, "fooey").WithCause(errors.New("cause"))
		require.Error(t, e)
		require.Equal(t, "fooey: cause", fmt.Sprintf("%v", e))
	})
	t.Run("+v", func(t *testing.T) {
		DefaultPackageName = "httperr"
		defer func() {
			DefaultPackageName = ""
		}()
		ln := lineNumber() + 1
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		out := fmt.Sprintf("%+v", e)
		lines := strings.Split(out, "\n")
		require.Len(t, lines, 3)
		require.Equal(t, "fooey", lines[0])
		require.Equal(t, "Stack:", lines[1])
		require.True(t, strings.HasPrefix(lines[2], "\tgithub.com/go-andiamo/httperr."))
		require.Contains(t, lines[2], ".TestError_Format.")
		require.True(t, strings.HasSuffix(lines[2], fmt.Sprintf(":%d", ln)))
	})
	t.Run("+v with cause", func(t *testing.T) {
		DefaultPackageName = "httperr"
		defer func() {
			DefaultPackageName = ""
		}()
		ln := lineNumber() + 1
		e := New(http.StatusBadRequest, "fooey").WithCause(errors.New("cause"))
		require.Error(t, e)
		out := fmt.Sprintf("%+v", e)
		lines := strings.Split(out, "\n")
		require.Len(t, lines, 3)
		require.Equal(t, "fooey: cause", lines[0])
		require.Equal(t, "Stack:", lines[1])
		require.True(t, strings.HasPrefix(lines[2], "\tgithub.com/go-andiamo/httperr."))
		require.Contains(t, lines[2], ".TestError_Format.")
		require.True(t, strings.HasSuffix(lines[2], fmt.Sprintf(":%d", ln)))
	})
	t.Run("s", func(t *testing.T) {
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		require.Equal(t, "fooey", fmt.Sprintf("%s", e))
	})
	t.Run("q", func(t *testing.T) {
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		require.Equal(t, `"fooey"`, fmt.Sprintf("%q", e))
	})
	t.Run("unknown", func(t *testing.T) {
		e := New(http.StatusBadRequest, "fooey")
		require.Error(t, e)
		require.Equal(t, "%!d(httperr.HttpError)", fmt.Sprintf("%d", e))
	})
}

func TestError_AddHeaders(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey")
	require.Error(t, e)
	require.Empty(t, e.Headers())

	_ = e.AddHeaders(map[string]string{"X-Foo": "bar"})
	require.Len(t, e.Headers(), 1)

	_ = e.AddHeaders(map[string]string{"X-Foo": "bar", "X-Foo2": "bar2"})
	require.Len(t, e.Headers(), 2)
}

func TestError_AddHeader(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey")
	require.Error(t, e)
	require.Empty(t, e.Headers())

	_ = e.AddHeader("X-Foo", "bar")
	require.Len(t, e.Headers(), 1)

	_ = e.AddHeader("X-Foo", "bar")
	require.Len(t, e.Headers(), 1)
}

func TestError_AddReasons(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey")
	require.Error(t, e)
	require.Empty(t, e.Reasons())

	_ = e.AddReasons("something", "something else")
	require.Len(t, e.Reasons(), 2)
}

func TestError_AddReason(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey")
	require.Error(t, e)
	require.Empty(t, e.Reasons())

	_ = e.AddReason("something")
	require.Len(t, e.Reasons(), 1)
	_ = e.AddReason("something else")
	require.Len(t, e.Reasons(), 2)
}

func TestError_Write(t *testing.T) {
	t.Run("with default writer", func(t *testing.T) {
		e := New(http.StatusBadRequest, "fooey").AddHeaders(map[string]string{"X-Foo": "bar"})
		w := httptest.NewRecorder()
		e.Write(w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		require.Equal(t, "bar", w.Header().Get("X-Foo"))
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "fooey", body[ptyError])
	})
	t.Run("no default writer", func(t *testing.T) {
		DefaultErrorWriter = nil
		defer func() {
			DefaultErrorWriter = &errorWriter{}
		}()
		e := New(http.StatusBadRequest, "fooey").AddHeaders(map[string]string{"X-Foo": "bar"})
		w := httptest.NewRecorder()
		e.Write(w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		require.Equal(t, "bar", w.Header().Get("X-Foo"))
	})
}

func TestError_MarshalJSON(t *testing.T) {
	e := New(http.StatusBadRequest, "fooey").
		WithCause(errors.New("cause")).
		AddReasons(struct {
			Reason string `json:"reason"`
		}{Reason: "bar"})
	t.Run("default", func(t *testing.T) {
		j, err := json.Marshal(e)
		require.NoError(t, err)
		fmt.Println(string(j))
		m := map[string]any{}
		require.NoError(t, json.Unmarshal(j, &m))
		require.Len(t, m, 2)
		require.Equal(t, "fooey", m[ptyError])
		require.Len(t, m[ptyReasons], 1)
	})
	t.Run("with cause", func(t *testing.T) {
		DefaultErrorWriterShowCause = true
		defer func() {
			DefaultErrorWriterShowCause = false
		}()
		j, err := json.Marshal(e)
		require.NoError(t, err)
		fmt.Println(string(j))
		m := map[string]any{}
		require.NoError(t, json.Unmarshal(j, &m))
		require.Len(t, m, 3)
		require.Equal(t, "fooey", m[ptyError])
		require.Len(t, m[ptyReasons], 1)
		require.Equal(t, "cause", m[ptyCause])
	})
	t.Run("with stack", func(t *testing.T) {
		DefaultErrorWriterShowStack = true
		DefaultPackageName = "httperr/"
		defer func() {
			DefaultErrorWriterShowStack = false
			DefaultPackageName = ""
		}()
		j, err := json.Marshal(e)
		require.NoError(t, err)
		fmt.Println(string(j))
		m := map[string]any{}
		require.NoError(t, json.Unmarshal(j, &m))
		require.Len(t, m, 3)
		require.Equal(t, "fooey", m[ptyError])
		require.Len(t, m[ptyReasons], 1)
		require.Len(t, m[ptyStack], 2)
	})
}

func TestPackageFromFunction(t *testing.T) {
	full, short, parts := packageFromFunction("github.com/go-andiamo/httperr.TestSomething.func2")
	require.Equal(t, "github.com/go-andiamo/httperr", full)
	require.Equal(t, "httperr", short)
	require.Equal(t, []string{"github.com", "go-andiamo", "httperr"}, parts)

	full, short, parts = packageFromFunction("httperr.TestSomething.func2")
	require.Equal(t, "httperr", full)
	require.Equal(t, "httperr", short)
	require.Equal(t, []string{"httperr"}, parts)
}

func lineNumber() int {
	const skip = 2
	pc := make([]uintptr, 1)
	n := runtime.Callers(skip, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	return frame.Line
}
