package errors

func E(code Code, detail Detail, err error) error {
	e := &Error{
		Code:   code,
		Detail: detail,
		frame:  Caller(1),
		err:    err,
	}
	return e
}
