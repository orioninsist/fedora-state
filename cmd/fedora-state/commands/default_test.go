package commands

import "testing"

func TestDefaultRunners(t *testing.T) {
	values := DefaultRunners()

	if len(values) == 0 {
		t.Fatal("expected runners")
	}

	if _, ok := values["plan"]; !ok {
		t.Fatal("expected plan runner")
	}
}
