// +build go1.13

package errors

import "errors"

type (
	Wrapper   = errors.Wrapper
	Formatter = errors.Formatter
	Printer   = errors.Printer
	Frame     = errors.Frame
)

var (
	As     = errors.As
	Is     = errors.Is
	New    = errors.New
	Caller = errors.Caller
	Unwrap = errors.Unwrap
	Opaque = errors.Opaque
)
