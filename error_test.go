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
		DBDelOp     Op = "db.delete"
		ServerDelOp Op = "server.delete"
	)

	e := E(DBDelOp, ENOTFOUND, "this is error!")
	e2 := E(ServerDelOp, e)

	if IsKind(ENOTFOUND, e2) == false {
		t.Fatalf("want: %v have: %v", true, false)
	}

	t.Logf("%v", e2)
	t.Logf("%+v", e2)
	//fmt.Printf("%+v", e2)

	//b, _ := json.Marshal(e2)
	//fmt.Println(string(b))

}
