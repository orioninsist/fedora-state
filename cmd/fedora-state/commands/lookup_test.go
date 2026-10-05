package commands

import "testing"

func TestExists(t *testing.T) {
	if !Exists("plan") {
		t.Fatal("expected plan command")
	}

	if Exists("unknown") {
		t.Fatal("did not expect unknown command")
	}
}
