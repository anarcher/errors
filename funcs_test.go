package errors

import "testing"

func TestBestError(t *testing.T) {
	const (
		notFound Kind = "notFound"
		found    Kind = "found"

		dbOp     Op = "db"
		serverOp Op = "server"
	)

	type check func(*Error, *testing.T)
	checks := func(cs ...check) []check { return cs }

	isKind := func(kind Kind) check {
		return func(err *Error, t *testing.T) {
			if err.Kind != kind {
				t.Fatalf("isKind: have: %v want: %v", err.Kind, kind)
			}
		}
	}
	isMessage := func(msg string) check {
		return func(err *Error, t *testing.T) {
			if err.Message != msg {
				t.Fatalf("isMsg: have: %v want: %v", err.Message, msg)
			}
		}
	}
	isOp := func(op Op) check {
		return func(err *Error, t *testing.T) {
			if err.Op != op {
				t.Fatalf("isOp: have: %v want: %v", err.Op, op)
			}
		}
	}

	tests := []struct {
		desc   string
		err    *Error
		checks []check
	}{
		{
			desc: "found kind",
			err: func() *Error {
				e1 := E(notFound, "msg1", dbOp)
				e2 := E(found, e1)
				return BestError(e2)
			}(),
			checks: checks(
				isKind(found),
				isMessage("msg1"),
				isOp(dbOp),
			),
		},
		{
			desc: "best error has a kind argument",
			err: func() *Error {
				e1 := E(notFound)
				e2 := E(e1, "msg1", serverOp)
				return BestError(e2, found, dbOp)
			}(),
			checks: checks(
				isKind(found),
				isMessage("msg1"),
				isOp(dbOp),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			for _, c := range test.checks {
				c(test.err, t)
			}
		})
	}
}
