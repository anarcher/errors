package errors

func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}

	e := &Error{}

	for _, arg := range args {
		switch arg := arg.(type) {
		case Kind:
			e.Kind = arg
		case Op:
			e.Op = arg
		case string:
			e.Message = arg
		case error:
			e.Err = arg
		}
	}

	e.frame = Caller(1)
	return e
}
