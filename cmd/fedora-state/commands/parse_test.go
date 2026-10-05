package commands

import "testing"

func TestResolve(t *testing.T) {
	if Resolve("plan") != "plan" {
		t.Fatal("expected plan")
	}

	if Resolve("unknown") != "help" {
		t.Fatal("expected help fallback")
	}
}
