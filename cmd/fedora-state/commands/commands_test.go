package commands

import "testing"

func TestAvailableCommands(t *testing.T) {
	values := Available()

	if len(values) == 0 {
		t.Fatal("expected commands")
	}
}
