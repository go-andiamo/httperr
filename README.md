# httperr
[![GoDoc](https://godoc.org/github.com/go-andiamo/httperr?status.svg)](https://pkg.go.dev/github.com/go-andiamo/httperr)
[![Latest Version](https://img.shields.io/github/v/tag/go-andiamo/httperr.svg?sort=semver&style=flat&label=version&color=blue)](https://github.com/go-andiamo/httperr/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-andiamo/httperr)](https://goreportcard.com/report/github.com/go-andiamo/httperr)

httperr is a Go package for HTTP/API errors with status codes and much more...

---

## Features

- Comprehensive helper functions for new errors (e.g. `NewBadRequestError()` and many more)
- Status code
- Errors with stack trace
- Cause, additional headers and reasons support
- Support for `errors.Unwrap`
- `Wrap` function - with pluggable support for wrapped error to status code resolution
- Configurable (and pluggable) error writer
- Pluggable formatting

---

## Installation

```bash
go get github.com/go-andiamo/httperr
```

