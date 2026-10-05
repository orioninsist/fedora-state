package commands

import "testing"

func TestRun(t *testing.T) {
	called := false

	err := Run(
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
		t.Fatal("runner not called")
	}
}
