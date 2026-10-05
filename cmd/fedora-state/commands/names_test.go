package commands

import "testing"

func TestCommandNames(t *testing.T) {
	if Plan == "" || Apply == "" || Manifest == "" {
		t.Fatal("missing command names")
	}
}
