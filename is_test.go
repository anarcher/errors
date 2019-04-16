package errors

import (
	"testing"
)

func TestIsKind(t *testing.T) {
	const (
		ENOTFOUND = Kind("notfound")
		EINVALID  = Kind("invalid")
	)

	const (
		DBDelOp     Op = "db.delete"
		ServerDelOp Op = "server.delete"
	)

	type check func(error, *testing.T)
	checks := func(cs ...check) []check { return cs }

	isKind := func(err error, t *testing.T) {
		if IsKind(err, ENOTFOUND) == false {
			t.Fatalf("IsKind: want: %v have: %v", true, false)
		}
	}

	tests := []struct {
		desc   string
		err    error
		checks []check
	}{
		{
			desc: "overlapping Kind",
			err:  E(ServerDelOp, E(DBDelOp, ENOTFOUND, "this is error")),
			checks: checks(
				isKind,
			),
		},
		{
			desc: "wrapping error only",
			err:  E(E(ServerDelOp, E(DBDelOp, ENOTFOUND, "this is err"))),
			checks: checks(
				isKind,
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
