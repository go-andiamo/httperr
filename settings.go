package httperr

import (
	"fmt"
	"runtime"
)

// PackageFilter is the interface used by DefaultPackageFilter
type PackageFilter interface {
	Include(packageName string) bool
}

// SetDefaultPackageFilter sets the DefaultPackageFilter with a filter for the specified package
func SetDefaultPackageFilter(pkg string) {
	DefaultPackageFilter = &packageFilter{
		packageName: pkg,
	}
}

type packageFilter struct {
	packageName string
}

var _ PackageFilter = (*packageFilter)(nil)

func (pf *packageFilter) Include(packageName string) bool {
	return pf.packageName == packageName
}

type FrameFormatter interface {
	StartLine() string
	FrameLine(frame runtime.Frame) string
}

type frameFormatter struct{}

var _ FrameFormatter = (*frameFormatter)(nil)

func (ff *frameFormatter) StartLine() string {
	return "\nStack:"
}

func (ff *frameFormatter) FrameLine(frame runtime.Frame) string {
	return fmt.Sprintf("\n\t%s:%d", frame.Function, frame.Line)
}

// ErrorStatusResolver is an interface for resolving status codes from errors
//
// and is the interface that DefaultErrorStatusResolver can be set to and used by Wrap
type ErrorStatusResolver interface {
	Resolve(err error, fallback int) int
}

// DefaultPackageFilter is the default package filter used to determine which packages
// are to be captured for the errors stack info
var DefaultPackageFilter PackageFilter

// DefaultPackageName is the default (short) package name to be used to check which packages
// are to be captured for the errors stack info
//
// if this var ends with "/", then any package that contains this name is included
var DefaultPackageName string

// MaxStackDepth is the maximum stack depth to capture
var MaxStackDepth uint = 16

// DefaultFrameFormatter is the formatter used to format call stack frames when formatting StackError
//
// If this is set to nil, no stack info is output when formatting StackError
var DefaultFrameFormatter FrameFormatter = &frameFormatter{}

// DefaultErrorStatusResolver is used by Wrap to determine the status codes for errors
var DefaultErrorStatusResolver ErrorStatusResolver
