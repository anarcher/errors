package errors

// ErrorKind returns the error which has Kind field
func ErrorKind(err error) *Error {
	e, ok := err.(*Error)
	if ok && e.Kind != "" {
		return e
	}
	if err := Unwrap(err); err != nil {
		return ErrorKind(err)
	}
	return nil
}

// ErrorMessage returns the message of an error which first met
func ErrorMessage(err error) *Error {
	e, ok := err.(*Error)
	if ok && e.Message != "" {
		return e
	}
	if err := Unwrap(err); err != nil {
		return ErrorMessage(err)
	}
	return nil

}

// ErrorOp returnr first errors which has Op in the chain of errors
func ErrorOp(err error) *Error {
	e, ok := err.(*Error)
	if ok && e.Op != "" {
		return e
	}
	if err := Unwrap(err); err != nil {
		return ErrorOp(err)
	}
	return nil
}

// BestError returns a new error with latest Kind, Op and Message
// of the chain of error
func BestError(err error, args ...interface{}) *Error {
	var e *Error
	if len(args) <= 0 {
		e = &Error{}
	} else {
		e = E(args...).(*Error)
	}

	if e.Kind == "" {
		if err := ErrorKind(err); err != nil {
			e.Kind = err.Kind
		}
	}

	if e.Message == "" {
		if err := ErrorMessage(err); err != nil {
			e.Message = err.Message
		}
	}

	if e.Op == "" {
		if err := ErrorOp(err); err != nil {
			e.Op = err.Op
		}
	}

	e.Err = err
	return e
}
