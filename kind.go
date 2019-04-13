package errors

// Is reports that err.Kind and kind are equal or not.
func IsKind(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if ok {
		if e.Kind == kind {
			return true
		}
	}

	if err := Unwrap(err); err != nil {
		return IsKind(kind, err)
	}

	return false
}

// AsKind returns Error value matched this kind.
func AsKind(kind Kind, err error) *Error {
	e, ok := err.(*Error)
	if ok {
		if e.Kind == kind {
			return e
		}
	}
	if err := Unwrap(err); err != nil {
		return AsKind(kind, err)
	}
	return nil
}

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
