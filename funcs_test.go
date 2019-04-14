package errors

import "testing"

func TestBestError(t *testing.T) {
	const (
		notFound Kind = "notFound"
		found    Kind = "found"
	)

	e1 := E(notFound, "message1")
	e2 := E(found, e1)

	bestErr := BestError(e2)

	if bestErr.Kind != found {
		t.Fatalf("kind: have: %v want: %v", bestErr.Kind, found)
	}

	if bestErr.Message != "message1" {
		t.Fatalf("message: have: %v want: %v", bestErr.Message, "message1")
	}

	t.Logf("bestErr: %+v", bestErr)
}
