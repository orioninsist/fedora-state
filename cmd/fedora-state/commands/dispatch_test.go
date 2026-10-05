package commands

import "testing"

func TestDispatch(t *testing.T) {
	called := false

	err := Dispatch(
		Plan,
		Context{},
		map[string]Handler{
			Plan: func(Context) error {
				called = true
				return nil
			},
			Help: func(Context) error {
				return nil
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("plan handler not called")
	}
}
