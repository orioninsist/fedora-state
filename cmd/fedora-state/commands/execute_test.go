package commands

import "testing"

func TestExecute(t *testing.T) {
	called := false

	err := Execute(
		"plan",
		map[string]Runner{
			"plan": func() error {
				called = true
				return nil
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("expected execution")
	}
}
