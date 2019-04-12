// +build !go1.13

package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

func (e *Error) Format(f fmt.State, c rune) { // implements fmt.Formatter
	xerrors.FormatError(e, f, c) // will call e.FormatError
}
