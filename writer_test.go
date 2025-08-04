package httperr

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultErrorWriter(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := NewBadRequestError("whoops")
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		require.Equal(t, applicationJson, w.Header().Get("Content-Type"))
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "whoops", body[ptyError])
	})
	t.Run("default (no show cause)", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := NewBadRequestError("whoops").WithCause(errors.New("something bad happened"))
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "whoops", body[ptyError])
	})
	t.Run("default (show cause)", func(t *testing.T) {
		DefaultErrorWriterShowCause = true
		defer func() {
			DefaultErrorWriterShowCause = false
		}()
		w := httptest.NewRecorder()
		e := NewBadRequestError("whoops").WithCause(errors.New("something bad happened"))
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 2)
		require.Equal(t, "whoops", body[ptyError])
		require.Equal(t, "something bad happened", body[ptyCause])
	})
	t.Run("default (show stack)", func(t *testing.T) {
		DefaultErrorWriterShowStack = true
		DefaultPackageName = "httperr"
		defer func() {
			DefaultErrorWriterShowStack = false
			DefaultPackageName = ""
		}()
		w := httptest.NewRecorder()
		ln := lineNumber() + 1
		e := NewBadRequestError("whoops")
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 2)
		require.Equal(t, "whoops", body[ptyError])
		stack, ok := body[prtStack].([]any)
		require.True(t, ok)
		require.Len(t, stack, 1)
		entry := stack[0].(string)
		require.Contains(t, entry, "github.com/go-andiamo/httperr")
		require.Contains(t, entry, ".TestDefaultErrorWriter.")
		require.Contains(t, entry, fmt.Sprintf(":%d", ln))
	})
	t.Run("default (with reason)", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := NewBadRequestError("whoops").AddReason(testReason{"foo", "too big"})
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 2)
		require.Equal(t, "whoops", body[ptyError])
		reasons, ok := body[ptyReasons].([]any)
		require.True(t, ok)
		require.Len(t, reasons, 1)
		reason, ok := reasons[0].(map[string]any)
		require.True(t, ok)
		require.Len(t, reason, 2)
		require.Equal(t, "foo", reason["property"])
		require.Equal(t, "too big", reason["reason"])
	})
	t.Run("default (with reasons)", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := NewBadRequestError("whoops").AddReasons(
			testReason{"foo", "too big"},
			testReason{"bar", "too small"},
		)
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 2)
		require.Equal(t, "whoops", body[ptyError])
		reasons, ok := body[ptyReasons].([]any)
		require.True(t, ok)
		require.Len(t, reasons, 2)
		reason, ok := reasons[0].(map[string]any)
		require.True(t, ok)
		require.Len(t, reason, 2)
		require.Equal(t, "foo", reason["property"])
		require.Equal(t, "too big", reason["reason"])
	})
	t.Run("default (with added headers)", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := NewBadRequestError("whoops").AddHeaders(map[string]string{
			"foo": "bar",
			"bar": "baz",
		})
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "whoops", body[ptyError])
		require.Equal(t, "bar", w.Result().Header.Get("foo"))
		require.Equal(t, "baz", w.Result().Header.Get("bar"))
	})
	t.Run("StatusError", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := &testStatusError{"whoops", http.StatusTeapot}
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusTeapot, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "whoops", body[ptyError])
	})
	t.Run("StatusError (empty message)", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := &testStatusError{"", http.StatusTeapot}
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusTeapot, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "I'm a teapot", body[ptyError])
	})
	t.Run("plain error", func(t *testing.T) {
		w := httptest.NewRecorder()
		e := errors.New("whoops")
		DefaultErrorWriter.WriteError(e, w)
		require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		body, err := unmarshalBody(w.Result().Body)
		require.NoError(t, err)
		require.Len(t, body, 1)
		require.Equal(t, "whoops", body[ptyError])
	})
}

type testReason struct {
	Pty    string `json:"property"`
	Reason string `json:"reason"`
}

type testStatusError struct {
	message string
	status  int
}

func (e *testStatusError) Error() string {
	return e.message
}

func (e *testStatusError) StatusCode() int {
	return e.status
}

var _ StatusError = (*testStatusError)(nil)

func unmarshalBody(body io.ReadCloser) (map[string]any, error) {
	result := make(map[string]any)
	err := json.NewDecoder(body).Decode(&result)
	return result, err
}
