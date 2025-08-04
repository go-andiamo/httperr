package httperr

import (
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	e := NewBadRequestError("")
	require.Equal(t, http.StatusBadRequest, e.StatusCode())
	require.Equal(t, "Bad Request", e.Error())
}

func TestNewBadRequestErrorf(t *testing.T) {
	e := NewBadRequestErrorf("something %d", 1)
	require.Equal(t, http.StatusBadRequest, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	e := NewUnauthorizedError("")
	require.Equal(t, http.StatusUnauthorized, e.StatusCode())
	require.Equal(t, "Unauthorized", e.Error())
}

func TestNewUnauthorizedErrorf(t *testing.T) {
	e := NewUnauthorizedErrorf("something %d", 1)
	require.Equal(t, http.StatusUnauthorized, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewPaymentRequiredError(t *testing.T) {
	e := NewPaymentRequiredError("")
	require.Equal(t, http.StatusPaymentRequired, e.StatusCode())
	require.Equal(t, "Payment Required", e.Error())
}

func TestNewPaymentRequiredErrorf(t *testing.T) {
	e := NewPaymentRequiredErrorf("something %d", 1)
	require.Equal(t, http.StatusPaymentRequired, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewForbiddenError(t *testing.T) {
	e := NewForbiddenError("")
	require.Equal(t, http.StatusForbidden, e.StatusCode())
	require.Equal(t, "Forbidden", e.Error())
}

func TestNewForbiddenErrorf(t *testing.T) {
	e := NewForbiddenErrorf("something %d", 1)
	require.Equal(t, http.StatusForbidden, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewNotFoundError(t *testing.T) {
	e := NewNotFoundError("")
	require.Equal(t, http.StatusNotFound, e.StatusCode())
	require.Equal(t, "Not Found", e.Error())
}

func TestNewNotFoundErrorf(t *testing.T) {
	e := NewNotFoundErrorf("something %d", 1)
	require.Equal(t, http.StatusNotFound, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewMethodNotAllowedError(t *testing.T) {
	e := NewMethodNotAllowedError("")
	require.Equal(t, http.StatusMethodNotAllowed, e.StatusCode())
	require.Equal(t, "Method Not Allowed", e.Error())
}

func TestNewMethodNotAllowedErrorf(t *testing.T) {
	e := NewMethodNotAllowedErrorf("something %d", 1)
	require.Equal(t, http.StatusMethodNotAllowed, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewNotAcceptableError(t *testing.T) {
	e := NewNotAcceptableError("")
	require.Equal(t, http.StatusNotAcceptable, e.StatusCode())
	require.Equal(t, "Not Acceptable", e.Error())
}

func TestNewNotAcceptableErrorf(t *testing.T) {
	e := NewNotAcceptableErrorf("something %d", 1)
	require.Equal(t, http.StatusNotAcceptable, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewProxyAuthRequiredError(t *testing.T) {
	e := NewProxyAuthRequiredError("")
	require.Equal(t, http.StatusProxyAuthRequired, e.StatusCode())
	require.Equal(t, "Proxy Authentication Required", e.Error())
}

func TestNewProxyAuthRequiredErrorf(t *testing.T) {
	e := NewProxyAuthRequiredErrorf("something %d", 1)
	require.Equal(t, http.StatusProxyAuthRequired, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewRequestTimeoutError(t *testing.T) {
	e := NewRequestTimeoutError("")
	require.Equal(t, http.StatusRequestTimeout, e.StatusCode())
	require.Equal(t, "Request Timeout", e.Error())
}

func TestNewRequestTimeoutErrorf(t *testing.T) {
	e := NewRequestTimeoutErrorf("something %d", 1)
	require.Equal(t, http.StatusRequestTimeout, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewConflictError(t *testing.T) {
	e := NewConflictError("")
	require.Equal(t, http.StatusConflict, e.StatusCode())
	require.Equal(t, "Conflict", e.Error())
}

func TestNewConflictErrorf(t *testing.T) {
	e := NewConflictErrorf("something %d", 1)
	require.Equal(t, http.StatusConflict, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewGoneError(t *testing.T) {
	e := NewGoneError("")
	require.Equal(t, http.StatusGone, e.StatusCode())
	require.Equal(t, "Gone", e.Error())
}

func TestNewGoneErrorf(t *testing.T) {
	e := NewGoneErrorf("something %d", 1)
	require.Equal(t, http.StatusGone, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewLengthRequiredError(t *testing.T) {
	e := NewLengthRequiredError("")
	require.Equal(t, http.StatusLengthRequired, e.StatusCode())
	require.Equal(t, "Length Required", e.Error())
}

func TestNewLengthRequiredErrorf(t *testing.T) {
	e := NewLengthRequiredErrorf("something %d", 1)
	require.Equal(t, http.StatusLengthRequired, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewPreconditionFailedError(t *testing.T) {
	e := NewPreconditionFailedError("")
	require.Equal(t, http.StatusPreconditionFailed, e.StatusCode())
	require.Equal(t, "Precondition Failed", e.Error())
}

func TestNewPreconditionFailedErrorf(t *testing.T) {
	e := NewPreconditionFailedErrorf("something %d", 1)
	require.Equal(t, http.StatusPreconditionFailed, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewRequestEntityTooLargeError(t *testing.T) {
	e := NewRequestEntityTooLargeError("")
	require.Equal(t, http.StatusRequestEntityTooLarge, e.StatusCode())
	require.Equal(t, "Request Entity Too Large", e.Error())
}

func TestNewRequestEntityTooLargeErrorf(t *testing.T) {
	e := NewRequestEntityTooLargeErrorf("something %d", 1)
	require.Equal(t, http.StatusRequestEntityTooLarge, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewRequestURITooLongError(t *testing.T) {
	e := NewRequestURITooLongError("")
	require.Equal(t, http.StatusRequestURITooLong, e.StatusCode())
	require.Equal(t, "Request URI Too Long", e.Error())
}

func TestNewRequestURITooLongErrorf(t *testing.T) {
	e := NewRequestURITooLongErrorf("something %d", 1)
	require.Equal(t, http.StatusRequestURITooLong, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewUnsupportedMediaTypeError(t *testing.T) {
	e := NewUnsupportedMediaTypeError("")
	require.Equal(t, http.StatusUnsupportedMediaType, e.StatusCode())
	require.Equal(t, "Unsupported Media Type", e.Error())
}

func TestNewUnsupportedMediaTypeErrorf(t *testing.T) {
	e := NewUnsupportedMediaTypeErrorf("something %d", 1)
	require.Equal(t, http.StatusUnsupportedMediaType, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewRequestedRangeNotSatisfiableError(t *testing.T) {
	e := NewRequestedRangeNotSatisfiableError("")
	require.Equal(t, http.StatusRequestedRangeNotSatisfiable, e.StatusCode())
	require.Equal(t, "Requested Range Not Satisfiable", e.Error())
}

func TestNewRequestedRangeNotSatisfiableErrorf(t *testing.T) {
	e := NewRequestedRangeNotSatisfiableErrorf("something %d", 1)
	require.Equal(t, http.StatusRequestedRangeNotSatisfiable, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewExpectationFailedError(t *testing.T) {
	e := NewExpectationFailedError("")
	require.Equal(t, http.StatusExpectationFailed, e.StatusCode())
	require.Equal(t, "Expectation Failed", e.Error())
}

func TestNewExpectationFailedErrorf(t *testing.T) {
	e := NewExpectationFailedErrorf("something %d", 1)
	require.Equal(t, http.StatusExpectationFailed, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewMisdirectedRequestError(t *testing.T) {
	e := NewMisdirectedRequestError("")
	require.Equal(t, http.StatusMisdirectedRequest, e.StatusCode())
	require.Equal(t, "Misdirected Request", e.Error())
}

func TestNewMisdirectedRequestErrorf(t *testing.T) {
	e := NewMisdirectedRequestErrorf("something %d", 1)
	require.Equal(t, http.StatusMisdirectedRequest, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewUnprocessableEntityError(t *testing.T) {
	e := NewUnprocessableEntityError("")
	require.Equal(t, http.StatusUnprocessableEntity, e.StatusCode())
	require.Equal(t, "Unprocessable Entity", e.Error())
}

func TestNewUnprocessableEntityErrorf(t *testing.T) {
	e := NewUnprocessableEntityErrorf("something %d", 1)
	require.Equal(t, http.StatusUnprocessableEntity, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewLockedError(t *testing.T) {
	e := NewLockedError("")
	require.Equal(t, http.StatusLocked, e.StatusCode())
	require.Equal(t, "Locked", e.Error())
}

func TestNewLockedErrorf(t *testing.T) {
	e := NewLockedErrorf("something %d", 1)
	require.Equal(t, http.StatusLocked, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewFailedDependencyError(t *testing.T) {
	e := NewFailedDependencyError("")
	require.Equal(t, http.StatusFailedDependency, e.StatusCode())
	require.Equal(t, "Failed Dependency", e.Error())
}

func TestNewFailedDependencyErrorf(t *testing.T) {
	e := NewFailedDependencyErrorf("something %d", 1)
	require.Equal(t, http.StatusFailedDependency, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewTooEarlyError(t *testing.T) {
	e := NewTooEarlyError("")
	require.Equal(t, http.StatusTooEarly, e.StatusCode())
	require.Equal(t, "Too Early", e.Error())
}

func TestNewTooEarlyErrorf(t *testing.T) {
	e := NewTooEarlyErrorf("something %d", 1)
	require.Equal(t, http.StatusTooEarly, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewUpgradeRequiredError(t *testing.T) {
	e := NewUpgradeRequiredError("")
	require.Equal(t, http.StatusUpgradeRequired, e.StatusCode())
	require.Equal(t, "Upgrade Required", e.Error())
}

func TestNewUpgradeRequiredErrorf(t *testing.T) {
	e := NewUpgradeRequiredErrorf("something %d", 1)
	require.Equal(t, http.StatusUpgradeRequired, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewPreconditionRequiredError(t *testing.T) {
	e := NewPreconditionRequiredError("")
	require.Equal(t, http.StatusPreconditionRequired, e.StatusCode())
	require.Equal(t, "Precondition Required", e.Error())
}

func TestNewPreconditionRequiredErrorf(t *testing.T) {
	e := NewPreconditionRequiredErrorf("something %d", 1)
	require.Equal(t, http.StatusPreconditionRequired, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewTooManyRequestsError(t *testing.T) {
	e := NewTooManyRequestsError("")
	require.Equal(t, http.StatusTooManyRequests, e.StatusCode())
	require.Equal(t, "Too Many Requests", e.Error())
}

func TestNewTooManyRequestsErrorf(t *testing.T) {
	e := NewTooManyRequestsErrorf("something %d", 1)
	require.Equal(t, http.StatusTooManyRequests, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewRequestHeaderFieldsTooLargeError(t *testing.T) {
	e := NewRequestHeaderFieldsTooLargeError("")
	require.Equal(t, http.StatusRequestHeaderFieldsTooLarge, e.StatusCode())
	require.Equal(t, "Request Header Fields Too Large", e.Error())
}

func TestNewRequestHeaderFieldsTooLargeErrorf(t *testing.T) {
	e := NewRequestHeaderFieldsTooLargeErrorf("something %d", 1)
	require.Equal(t, http.StatusRequestHeaderFieldsTooLarge, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewUnavailableForLegalReasonsError(t *testing.T) {
	e := NewUnavailableForLegalReasonsError("")
	require.Equal(t, http.StatusUnavailableForLegalReasons, e.StatusCode())
	require.Equal(t, "Unavailable For Legal Reasons", e.Error())
}

func TestNewUnavailableForLegalReasonsErrorf(t *testing.T) {
	e := NewUnavailableForLegalReasonsErrorf("something %d", 1)
	require.Equal(t, http.StatusUnavailableForLegalReasons, e.StatusCode())
	require.Equal(t, "something 1", e.Error())
}

func TestNewInternalServerError(t *testing.T) {
	e := NewInternalServerError("", errors.New("cause"))
	require.Equal(t, http.StatusInternalServerError, e.StatusCode())
	require.Equal(t, "Internal Server Error", e.Error())
	require.Error(t, errors.Unwrap(e))
}

func TestNewNotImplementedError(t *testing.T) {
	e := NewNotImplementedError("")
	require.Equal(t, http.StatusNotImplemented, e.StatusCode())
	require.Equal(t, "Not Implemented", e.Error())
}

func TestNewBadGatewayError(t *testing.T) {
	e := NewBadGatewayError("", errors.New("cause"))
	require.Equal(t, http.StatusBadGateway, e.StatusCode())
	require.Equal(t, "Bad Gateway", e.Error())
	require.Error(t, errors.Unwrap(e))
}

func TestNewServiceUnavailableError(t *testing.T) {
	e := NewServiceUnavailableError("", errors.New("cause"))
	require.Equal(t, http.StatusServiceUnavailable, e.StatusCode())
	require.Equal(t, "Service Unavailable", e.Error())
	require.Error(t, errors.Unwrap(e))
}

func TestNewGatewayTimeoutError(t *testing.T) {
	e := NewGatewayTimeoutError("")
	require.Equal(t, http.StatusGatewayTimeout, e.StatusCode())
	require.Equal(t, "Gateway Timeout", e.Error())
}

func TestNewHTTPVersionNotSupportedError(t *testing.T) {
	e := NewHTTPVersionNotSupportedError("")
	require.Equal(t, http.StatusHTTPVersionNotSupported, e.StatusCode())
	require.Equal(t, "HTTP Version Not Supported", e.Error())
}

func TestNewVariantAlsoNegotiatesError(t *testing.T) {
	e := NewVariantAlsoNegotiatesError("")
	require.Equal(t, http.StatusVariantAlsoNegotiates, e.StatusCode())
	require.Equal(t, "Variant Also Negotiates", e.Error())
}

func TestNewInsufficientStorageError(t *testing.T) {
	e := NewInsufficientStorageError("", errors.New("cause"))
	require.Equal(t, http.StatusInsufficientStorage, e.StatusCode())
	require.Equal(t, "Insufficient Storage", e.Error())
	require.Error(t, errors.Unwrap(e))
}

func TestNewLoopDetectedError(t *testing.T) {
	e := NewLoopDetectedError("", errors.New("cause"))
	require.Equal(t, http.StatusLoopDetected, e.StatusCode())
	require.Equal(t, "Loop Detected", e.Error())
	require.Error(t, errors.Unwrap(e))
}

func TestNewNotExtendedError(t *testing.T) {
	e := NewNotExtendedError("")
	require.Equal(t, http.StatusNotExtended, e.StatusCode())
	require.Equal(t, "Not Extended", e.Error())
}

func TestNewNetworkAuthenticationRequiredError(t *testing.T) {
	e := NewNetworkAuthRequiredError("")
	require.Equal(t, http.StatusNetworkAuthenticationRequired, e.StatusCode())
	require.Equal(t, "Network Authentication Required", e.Error())
}

func TestNewMultipleChoicesError(t *testing.T) {
	e := NewMultipleChoicesError("", "")
	require.Equal(t, http.StatusMultipleChoices, e.StatusCode())
	require.Equal(t, "Multiple Choices", e.Error())
	require.Empty(t, e.Headers())

	e = NewMultipleChoicesError("", "somewhere")
	require.Equal(t, http.StatusMultipleChoices, e.StatusCode())
	require.Equal(t, "Multiple Choices", e.Error())
	require.NotEmpty(t, e.Headers())
	require.Equal(t, "somewhere", e.Headers()["Location"])
}

func TestNewMovedPermanentlyError(t *testing.T) {
	e := NewMovedPermanentlyError("", "")
	require.Equal(t, http.StatusMovedPermanently, e.StatusCode())
	require.Equal(t, "Moved Permanently", e.Error())
	require.Empty(t, e.Headers())

	e = NewMovedPermanentlyError("", "somewhere")
	require.Equal(t, http.StatusMovedPermanently, e.StatusCode())
	require.Equal(t, "Moved Permanently", e.Error())
	require.NotEmpty(t, e.Headers())
	require.Equal(t, "somewhere", e.Headers()["Location"])
}

func TestNewFoundError(t *testing.T) {
	e := NewFoundError("", "")
	require.Equal(t, http.StatusFound, e.StatusCode())
	require.Equal(t, "Found", e.Error())
	require.Empty(t, e.Headers())

	e = NewFoundError("", "somewhere")
	require.Equal(t, http.StatusFound, e.StatusCode())
	require.Equal(t, "Found", e.Error())
	require.NotEmpty(t, e.Headers())
	require.Equal(t, "somewhere", e.Headers()["Location"])
}

func TestNewSeeOtherError(t *testing.T) {
	e := NewSeeOtherError("", "")
	require.Equal(t, http.StatusSeeOther, e.StatusCode())
	require.Equal(t, "See Other", e.Error())
	require.Empty(t, e.Headers())

	e = NewSeeOtherError("", "somewhere")
	require.Equal(t, http.StatusSeeOther, e.StatusCode())
	require.Equal(t, "See Other", e.Error())
	require.NotEmpty(t, e.Headers())
	require.Equal(t, "somewhere", e.Headers()["Location"])
}

func TestNewNotModifiedError(t *testing.T) {
	e := NewNotModifiedError("")
	require.Equal(t, http.StatusNotModified, e.StatusCode())
	require.Equal(t, "Not Modified", e.Error())
}

func TestNewTemporaryRedirectError(t *testing.T) {
	e := NewTemporaryRedirectError("", "")
	require.Equal(t, http.StatusTemporaryRedirect, e.StatusCode())
	require.Equal(t, "Temporary Redirect", e.Error())
	require.Empty(t, e.Headers())

	e = NewTemporaryRedirectError("", "somewhere")
	require.Equal(t, http.StatusTemporaryRedirect, e.StatusCode())
	require.Equal(t, "Temporary Redirect", e.Error())
	require.NotEmpty(t, e.Headers())
	require.Equal(t, "somewhere", e.Headers()["Location"])
}

func TestNewPermanentRedirectError(t *testing.T) {
	e := NewPermanentRedirectError("", "")
	require.Equal(t, http.StatusPermanentRedirect, e.StatusCode())
	require.Equal(t, "Permanent Redirect", e.Error())
	require.Empty(t, e.Headers())

	e = NewPermanentRedirectError("", "somewhere")
	require.Equal(t, http.StatusPermanentRedirect, e.StatusCode())
	require.Equal(t, "Permanent Redirect", e.Error())
	require.NotEmpty(t, e.Headers())
	require.Equal(t, "somewhere", e.Headers()["Location"])
}
