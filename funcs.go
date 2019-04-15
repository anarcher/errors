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

// BestError returns a new error with latest Kind and Message
// of the chain of error
func BestError(err error, args ...interface{}) *Error {
	var e *Error
	if len(args) <= 0 {
		e = &Error{}
	} else {
		e = E(args...).(*Error)
	}

	if err := ErrorKind(err); err != nil {
		e.Kind = err.Kind
	}

	if err := ErrorMessage(err); err != nil {
		e.Message = err.Message
	}

	e.Err = err
	return e
}
