package commands

import "testing"

func TestCommandFromArgs(t *testing.T) {
	if CommandFromArgs([]string{"plan"}) != "plan" {
		t.Fatal("expected plan")
	}

	if CommandFromArgs(nil) != "help" {
		t.Fatal("expected help")
	}
}
