package errors

import "testing"

type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return "path: " + e.Path
}

func TestError(t *testing.T) {

	eof := New("EOF")
	e := E(eof, &PathError{"/home"}, nil)

	if Is(e, eof) != true {
		t.Errorf("want: %v have: %v", true, Is(e, eof))
	}

	var pathErr *PathError
	if have := As(e, &pathErr); have == false {
		t.Errorf("want: %v have: %v", true, have)
	}

	t.Logf("%+v", e)
}
