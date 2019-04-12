// +build !go1.13

package errors

import "golang.org/x/xerrors"

type (
	Wrapper   = xerrors.Wrapper
	Formatter = xerrors.Formatter
	Printer   = xerrors.Printer
	Frame     = xerrors.Frame
)

var (
	As     = xerrors.As
	Is     = xerrors.Is
	New    = xerrors.New
	Caller = xerrors.Caller
	Unwrap = xerrors.Unwrap
	Opaque = xerrors.Opaque
)
