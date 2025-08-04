package httperr

import (
	"fmt"
	"net/http"
)

// NewBadRequestError creates a new 400 Bad Request error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-400-bad-request
func NewBadRequestError(msg string) HttpError {
	return newError(http.StatusBadRequest, msg, nil, getStackInfo())
}

// NewBadRequestErrorf creates a new 400 Bad Request error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-400-bad-request
func NewBadRequestErrorf(format string, a ...any) HttpError {
	return newError(http.StatusBadRequest, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewUnauthorizedError creates a new 401 Unauthorized error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-401-unauthorized
func NewUnauthorizedError(msg string) HttpError {
	return newError(http.StatusUnauthorized, msg, nil, getStackInfo())
}

// NewUnauthorizedErrorf creates a new 401 Unauthorized error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-401-unauthorized
func NewUnauthorizedErrorf(format string, a ...any) HttpError {
	return newError(http.StatusUnauthorized, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewPaymentRequiredError creates a new 402 Payment Required error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-402-payment-required
func NewPaymentRequiredError(msg string) HttpError {
	return newError(http.StatusPaymentRequired, msg, nil, getStackInfo())
}

// NewPaymentRequiredErrorf creates a new 402 Payment Required error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-402-payment-required
func NewPaymentRequiredErrorf(format string, a ...any) HttpError {
	return newError(http.StatusPaymentRequired, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewForbiddenError creates a new 403 Forbidden error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-403-forbidden
func NewForbiddenError(msg string) HttpError {
	return newError(http.StatusForbidden, msg, nil, getStackInfo())
}

// NewForbiddenErrorf creates a new 403 Forbidden error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-403-forbidden
func NewForbiddenErrorf(format string, a ...any) HttpError {
	return newError(http.StatusForbidden, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewNotFoundError creates a new 404 Not Found error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-404-not-found
func NewNotFoundError(msg string) HttpError {
	return newError(http.StatusNotFound, msg, nil, getStackInfo())
}

// NewNotFoundErrorf creates a new 404 Not Found error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-404-not-found
func NewNotFoundErrorf(format string, a ...any) HttpError {
	return newError(http.StatusNotFound, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewMethodNotAllowedError creates a new 405 Method Not Allowed error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-405-method-not-allowed
func NewMethodNotAllowedError(msg string) HttpError {
	return newError(http.StatusMethodNotAllowed, msg, nil, getStackInfo())
}

// NewMethodNotAllowedErrorf creates a new 405 Method Not Allowed error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-405-method-not-allowed
func NewMethodNotAllowedErrorf(format string, a ...any) HttpError {
	return newError(http.StatusMethodNotAllowed, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewNotAcceptableError creates a new 406 Not Acceptable error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-406-not-acceptable
func NewNotAcceptableError(msg string) HttpError {
	return newError(http.StatusNotAcceptable, msg, nil, getStackInfo())
}

// NewNotAcceptableErrorf creates a new 406 Not Acceptable error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-406-not-acceptable
func NewNotAcceptableErrorf(format string, a ...any) HttpError {
	return newError(http.StatusNotAcceptable, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewProxyAuthRequiredError creates a new 407 Proxy Authentication Required error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-407-proxy-authentication-re
func NewProxyAuthRequiredError(msg string) HttpError {
	return newError(http.StatusProxyAuthRequired, msg, nil, getStackInfo())
}

// NewProxyAuthRequiredErrorf creates a new 407 Proxy Authentication Required error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-407-proxy-authentication-re
func NewProxyAuthRequiredErrorf(format string, a ...any) HttpError {
	return newError(http.StatusProxyAuthRequired, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewRequestTimeoutError creates a new 408 Request Timeout error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-408-request-timeout
func NewRequestTimeoutError(msg string) HttpError {
	return newError(http.StatusRequestTimeout, msg, nil, getStackInfo())
}

// NewRequestTimeoutErrorf creates a new 408 Request Timeout error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-408-request-timeout
func NewRequestTimeoutErrorf(format string, a ...any) HttpError {
	return newError(http.StatusRequestTimeout, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewConflictError creates a new 409 Conflict error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-409-conflict
func NewConflictError(msg string) HttpError {
	return newError(http.StatusConflict, msg, nil, getStackInfo())
}

// NewConflictErrorf creates a new 409 Conflict error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-409-conflict
func NewConflictErrorf(format string, a ...any) HttpError {
	return newError(http.StatusConflict, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewGoneError creates a new 410 Gone error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-410-gone
func NewGoneError(msg string) HttpError {
	return newError(http.StatusGone, msg, nil, getStackInfo())
}

// NewGoneErrorf creates a new 410 Gone error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-410-gone
func NewGoneErrorf(format string, a ...any) HttpError {
	return newError(http.StatusGone, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewLengthRequiredError creates a new 411 Length Required error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-411-length-required
func NewLengthRequiredError(msg string) HttpError {
	return newError(http.StatusLengthRequired, msg, nil, getStackInfo())
}

// NewLengthRequiredErrorf creates a new 411 Length Required error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-411-length-required
func NewLengthRequiredErrorf(format string, a ...any) HttpError {
	return newError(http.StatusLengthRequired, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewPreconditionFailedError creates a new 412 Precondition Failed error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-412-precondition-failed
func NewPreconditionFailedError(msg string) HttpError {
	return newError(http.StatusPreconditionFailed, msg, nil, getStackInfo())
}

// NewPreconditionFailedErrorf creates a new 412 Precondition Failed error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-412-precondition-failed
func NewPreconditionFailedErrorf(format string, a ...any) HttpError {
	return newError(http.StatusPreconditionFailed, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewRequestEntityTooLargeError creates a new 413 Request Entity Too Large error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-413-content-too-large
func NewRequestEntityTooLargeError(msg string) HttpError {
	return newError(http.StatusRequestEntityTooLarge, msg, nil, getStackInfo())
}

// NewRequestEntityTooLargeErrorf creates a new 413 Request Entity Too Large error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-413-content-too-large
func NewRequestEntityTooLargeErrorf(format string, a ...any) HttpError {
	return newError(http.StatusRequestEntityTooLarge, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewRequestURITooLongError creates a new 414 Request URI Too Long error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-414-uri-too-long
func NewRequestURITooLongError(msg string) HttpError {
	return newError(http.StatusRequestURITooLong, msg, nil, getStackInfo())
}

// NewRequestURITooLongErrorf creates a new 414 Request URI Too Long error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-414-uri-too-long
func NewRequestURITooLongErrorf(format string, a ...any) HttpError {
	return newError(http.StatusRequestURITooLong, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewUnsupportedMediaTypeError creates a new 415 Unsupported Media Type error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-415-unsupported-media-type
func NewUnsupportedMediaTypeError(msg string) HttpError {
	return newError(http.StatusUnsupportedMediaType, msg, nil, getStackInfo())
}

// NewUnsupportedMediaTypeErrorf creates a new 415 Unsupported Media Type error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-415-unsupported-media-type
func NewUnsupportedMediaTypeErrorf(format string, a ...any) HttpError {
	return newError(http.StatusUnsupportedMediaType, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewRequestedRangeNotSatisfiableError creates a new 416 Requested Range Not Satisfiable error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-416-range-not-satisfiable
func NewRequestedRangeNotSatisfiableError(msg string) HttpError {
	return newError(http.StatusRequestedRangeNotSatisfiable, msg, nil, getStackInfo())
}

// NewRequestedRangeNotSatisfiableErrorf creates a new 416 Requested Range Not Satisfiable error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-416-range-not-satisfiable
func NewRequestedRangeNotSatisfiableErrorf(format string, a ...any) HttpError {
	return newError(http.StatusRequestedRangeNotSatisfiable, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewExpectationFailedError creates a new 417 Expectation Failed error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-417-expectation-failed
func NewExpectationFailedError(msg string) HttpError {
	return newError(http.StatusExpectationFailed, msg, nil, getStackInfo())
}

// NewExpectationFailedErrorf creates a new 417 Expectation Failed error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-417-expectation-failed
func NewExpectationFailedErrorf(format string, a ...any) HttpError {
	return newError(http.StatusExpectationFailed, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewMisdirectedRequestError creates a new 421 Misdirected Request error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-421-misdirected-request
func NewMisdirectedRequestError(msg string) HttpError {
	return newError(http.StatusMisdirectedRequest, msg, nil, getStackInfo())
}

// NewMisdirectedRequestErrorf creates a new 421 Misdirected Request error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-421-misdirected-request
func NewMisdirectedRequestErrorf(format string, a ...any) HttpError {
	return newError(http.StatusMisdirectedRequest, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewUnprocessableEntityError creates a new 422 Unprocessable Entity error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-422-unprocessable-content
func NewUnprocessableEntityError(msg string) HttpError {
	return newError(http.StatusUnprocessableEntity, msg, nil, getStackInfo())
}

// NewUnprocessableEntityErrorf creates a new 422 Unprocessable Entity error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-422-unprocessable-content
func NewUnprocessableEntityErrorf(format string, a ...any) HttpError {
	return newError(http.StatusUnprocessableEntity, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewLockedError creates a new 423 Locked error
//
// see https://datatracker.ietf.org/doc/html/rfc4918#section-11.3
func NewLockedError(msg string) HttpError {
	return newError(http.StatusLocked, msg, nil, getStackInfo())
}

// NewLockedErrorf creates a new 423 Locked error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc4918#section-11.3
func NewLockedErrorf(format string, a ...any) HttpError {
	return newError(http.StatusLocked, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewFailedDependencyError creates a new 424 Failed Dependency error
//
// see https://datatracker.ietf.org/doc/html/rfc4918#section-11.4
func NewFailedDependencyError(msg string) HttpError {
	return newError(http.StatusFailedDependency, msg, nil, getStackInfo())
}

// NewFailedDependencyErrorf creates a new 424 Failed Dependency error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc4918#section-11.4
func NewFailedDependencyErrorf(format string, a ...any) HttpError {
	return newError(http.StatusFailedDependency, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewTooEarlyError creates a new 425 Too Early error
//
// see https://datatracker.ietf.org/doc/html/rfc8470#section-5.2
func NewTooEarlyError(msg string) HttpError {
	return newError(http.StatusTooEarly, msg, nil, getStackInfo())
}

// NewTooEarlyErrorf creates a new 425 Too Early error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc8470#section-5.2
func NewTooEarlyErrorf(format string, a ...any) HttpError {
	return newError(http.StatusTooEarly, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewUpgradeRequiredError creates a new 426 Upgrade Required error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-426-upgrade-required
func NewUpgradeRequiredError(msg string) HttpError {
	return newError(http.StatusUpgradeRequired, msg, nil, getStackInfo())
}

// NewUpgradeRequiredErrorf creates a new 426 Upgrade Required error with a formatted message
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-426-upgrade-required
func NewUpgradeRequiredErrorf(format string, a ...any) HttpError {
	return newError(http.StatusUpgradeRequired, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewPreconditionRequiredError creates a new 428 Precondition Required error
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-3
func NewPreconditionRequiredError(msg string) HttpError {
	return newError(http.StatusPreconditionRequired, msg, nil, getStackInfo())
}

// NewPreconditionRequiredErrorf creates a new 428 Precondition Required error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-3
func NewPreconditionRequiredErrorf(format string, a ...any) HttpError {
	return newError(http.StatusPreconditionRequired, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewTooManyRequestsError creates a new 429 Too Many Requests error
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-4
func NewTooManyRequestsError(msg string) HttpError {
	return newError(http.StatusTooManyRequests, msg, nil, getStackInfo())
}

// NewTooManyRequestsErrorf creates a new 429 Too Many Requests error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-4
func NewTooManyRequestsErrorf(format string, a ...any) HttpError {
	return newError(http.StatusTooManyRequests, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewRequestHeaderFieldsTooLargeError creates a new 431 Request Header Fields Too Large error
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-5
func NewRequestHeaderFieldsTooLargeError(msg string) HttpError {
	return newError(http.StatusRequestHeaderFieldsTooLarge, msg, nil, getStackInfo())
}

// NewRequestHeaderFieldsTooLargeErrorf creates a new 431 Request Header Fields Too Large error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-5
func NewRequestHeaderFieldsTooLargeErrorf(format string, a ...any) HttpError {
	return newError(http.StatusRequestHeaderFieldsTooLarge, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewUnavailableForLegalReasonsError creates a new 451 Unavailable For Legal Reasons error
//
// see https://datatracker.ietf.org/doc/html/rfc7725#section-3
func NewUnavailableForLegalReasonsError(msg string) HttpError {
	return newError(http.StatusUnavailableForLegalReasons, msg, nil, getStackInfo())
}

// NewUnavailableForLegalReasonsErrorf creates a new 451 Unavailable For Legal Reasons error with a formatted message
//
// see https://datatracker.ietf.org/doc/html/rfc7725#section-3
func NewUnavailableForLegalReasonsErrorf(format string, a ...any) HttpError {
	return newError(http.StatusUnavailableForLegalReasons, fmt.Sprintf(format, a...), nil, getStackInfo())
}

// NewInternalServerError creates a new 500 Internal Server error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-500-internal-server-error
func NewInternalServerError(msg string, cause error) HttpError {
	return newError(http.StatusInternalServerError, msg, cause, getStackInfo())
}

// NewNotImplementedError creates a new 501 Not Implemented error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-501-not-implemented
func NewNotImplementedError(msg string) HttpError {
	return newError(http.StatusNotImplemented, msg, nil, getStackInfo())
}

// NewBadGatewayError creates a new 502 Bad Gateway error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-502-bad-gateway
func NewBadGatewayError(msg string, cause error) HttpError {
	return newError(http.StatusBadGateway, msg, cause, getStackInfo())
}

// NewServiceUnavailableError creates a new 503 Service Unavailable error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-503-service-unavailable
func NewServiceUnavailableError(msg string, cause error) HttpError {
	return newError(http.StatusServiceUnavailable, msg, cause, getStackInfo())
}

// NewGatewayTimeoutError creates a new 504 Gateway Timeout error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-504-gateway-timeout
func NewGatewayTimeoutError(msg string) HttpError {
	return newError(http.StatusGatewayTimeout, msg, nil, getStackInfo())
}

// NewHTTPVersionNotSupportedError creates a new 505 HTTP Version Not Supported error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-505-http-version-not-suppor
func NewHTTPVersionNotSupportedError(msg string) HttpError {
	return newError(http.StatusHTTPVersionNotSupported, msg, nil, getStackInfo())
}

// NewVariantAlsoNegotiatesError creates a new 506 Variant Also Negotiates error
//
// see https://datatracker.ietf.org/doc/html/rfc2295#section-8.1
func NewVariantAlsoNegotiatesError(msg string) HttpError {
	return newError(http.StatusVariantAlsoNegotiates, msg, nil, getStackInfo())
}

// NewInsufficientStorageError creates a new 507 Insufficient Storage error
//
// see https://datatracker.ietf.org/doc/html/rfc4918#section-11.5
func NewInsufficientStorageError(msg string, cause error) HttpError {
	return newError(http.StatusInsufficientStorage, msg, cause, getStackInfo())
}

// NewLoopDetectedError creates a new 508 Loop Detected error
//
// see https://www.rfc-editor.org/rfc/rfc5842.html#section-7.2
func NewLoopDetectedError(msg string, cause error) HttpError {
	return newError(http.StatusLoopDetected, msg, cause, getStackInfo())
}

// NewNotExtendedError creates a new 510 Not Extended error
//
// see https://datatracker.ietf.org/doc/html/rfc2774#section-7
func NewNotExtendedError(msg string) HttpError {
	return newError(http.StatusNotExtended, msg, nil, getStackInfo())
}

// NewNetworkAuthRequiredError creates a new 511 Network Authentication Required error
//
// see https://datatracker.ietf.org/doc/html/rfc6585#section-6
func NewNetworkAuthRequiredError(msg string) HttpError {
	return newError(http.StatusNetworkAuthenticationRequired, msg, nil, getStackInfo())
}

const hdrLocation = "Location"

// NewMultipleChoicesError creates a new 300 Multiple Choices error
//
// if the location arg is a non-empty string, a `Location` header is added
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-300-multiple-choices
func NewMultipleChoicesError(msg string, location string) HttpError {
	if location == "" {
		return newError(http.StatusMultipleChoices, msg, nil, getStackInfo())
	}
	return newError(http.StatusMultipleChoices, msg, nil, getStackInfo()).AddHeader(hdrLocation, location)
}

// NewMovedPermanentlyError creates a new 301 Moved Permanently error
//
// if the location arg is a non-empty string, a `Location` header is added
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-301-moved-permanently
func NewMovedPermanentlyError(msg string, location string) HttpError {
	if location == "" {
		return newError(http.StatusMovedPermanently, msg, nil, getStackInfo())
	}
	return newError(http.StatusMovedPermanently, msg, nil, getStackInfo()).AddHeader(hdrLocation, location)
}

// NewFoundError creates a new 302 Found error
//
// if the location arg is a non-empty string, a `Location` header is added
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-302-found
func NewFoundError(msg string, location string) HttpError {
	if location == "" {
		return newError(http.StatusFound, msg, nil, getStackInfo())
	}
	return newError(http.StatusFound, msg, nil, getStackInfo()).AddHeader(hdrLocation, location)
}

// NewSeeOtherError creates a new 303 See Other error
//
// if the location arg is a non-empty string, a `Location` header is added
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-302-found
func NewSeeOtherError(msg string, location string) HttpError {
	if location == "" {
		return newError(http.StatusSeeOther, msg, nil, getStackInfo())
	}
	return newError(http.StatusSeeOther, msg, nil, getStackInfo()).AddHeader(hdrLocation, location)
}

// NewNotModifiedError creates a new 304 Not Modified error
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-304-not-modified
func NewNotModifiedError(msg string) HttpError {
	return newError(http.StatusNotModified, msg, nil, getStackInfo())
}

// NewTemporaryRedirectError creates a new 307 Temporary Redirect error
//
// if the location arg is a non-empty string, a `Location` header is added
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-307-temporary-redirect
func NewTemporaryRedirectError(msg string, location string) HttpError {
	if location == "" {
		return newError(http.StatusTemporaryRedirect, msg, nil, getStackInfo())
	}
	return newError(http.StatusTemporaryRedirect, msg, nil, getStackInfo()).AddHeader(hdrLocation, location)
}

// NewPermanentRedirectError creates a new 308 Permanent Redirect error
//
// if the location arg is a non-empty string, a `Location` header is added
//
// see https://www.rfc-editor.org/rfc/rfc9110.html#name-308-permanent-redirect
func NewPermanentRedirectError(msg string, location string) HttpError {
	if location == "" {
		return newError(http.StatusPermanentRedirect, msg, nil, getStackInfo())
	}
	return newError(http.StatusPermanentRedirect, msg, nil, getStackInfo()).AddHeader(hdrLocation, location)
}
