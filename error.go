package errors

type Code interface {
	error
}

type Detail interface {
	error
}

type Error struct {
	Code   Code
	Detail Detail
	err    error
	frame  Frame
}

func (e *Error) Error() string {
	if e.Code == nil {
		return ""
	}
	return e.Code.Error()
}

func (e *Error) FormatError(p Printer) error {
	p.Print(e.Code)
	e.frame.Format(p)
	if p.Detail() && e.Detail != nil {
		p.Print(e.Detail.Error())
	}
	return e.err
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Is(err error) bool {
	return Is(e.Code, err)
}

func (e *Error) As(target interface{}) bool {
	return As(e.Detail, target)
}
