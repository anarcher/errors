package errors

import (
	"fmt"
	"strings"
)

// Kind defines the kind of error
type Kind string

// Op describes an operation
type Op string

// Error defines a application error
type Error struct {

	// Kind is the kind of this error
	Kind Kind `json:"kind"`

	// Human-readable message
	Message string `json:"message"`

	// Logical operation. usually the name of method.
	Op Op `json:"op"`

	// Nested error
	Err error `json:"error"`

	frame Frame
}

func (e *Error) Error() string {
	var b strings.Builder

	if e.Kind != "" {
		fmt.Fprintf(&b, "[%s]", e.Kind)
	}

	if e.Op != "" {
		if e.Kind != "" {
			b.WriteString(": ")
		}
		fmt.Fprintf(&b, "%s", e.Op)
	}

	return b.String()
}

func (e *Error) FormatError(p Printer) error {
	p.Print(e.Error())
	if p.Detail() {
		p.Print(e.Message)
	}
	e.frame.Format(p)
	return e.Unwrap()
}

func (e *Error) Unwrap() error {
	return e.Err
}
