package commands

import "testing"

func TestHandlerMap(t *testing.T) {
	values := HandlerMap()

	if _, ok := values[Plan]; !ok {
		t.Fatal("expected plan handler")
	}
}
