package errors

import "fmt"

func Match(err error, args ...interface{}) *Error {
	e, ok := err.(*Error)
	if !ok {
		if err := Unwrap(err); err != nil {
			return Match(err, args...)
		}
	}

	var matched int
	for _, arg := range args {
		switch arg.(type) {
		case Kind:
			if e.Kind == arg {
				matched++
			}
		case Op:
			if e.Op == arg {
				matched++
			}
		case string:
			if e.Message == arg {
				matched++
			}
		default:
			panic(fmt.Sprintf("match with unsupported arguments: %v<%T>", arg, arg))
		}
	}
	if len(args) == matched {
		return e
	}

	if err := Unwrap(err); err != nil {
		return Match(err, args...)
	}

	return nil
}
