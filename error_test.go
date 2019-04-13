package errors

import (
	"testing"
)

func TestError(t *testing.T) {
	const (
		ENOTFOUND = Kind("notfound")
		EINVALID  = Kind("invalid")
	)

	const (
		op1 Op = "db.delete"
		op2 Op = "server.delete"
	)

	e := E(op1, ENOTFOUND, "this is error!")
	e2 := E(op2, e)

	if IsKind(ENOTFOUND, e2) == false {
		t.Fatalf("want: %v have: %v", true, false)
	}

	t.Logf("%v", e2)
	t.Logf("%+v", e2)
	//fmt.Printf("%+v", e2)

	//b, _ := json.Marshal(e2)
	//fmt.Println(string(b))

}
