package plan

import "testing"

func TestEmptyPlan(t *testing.T) {
	got := Empty()

	if len(got.Actions) != 0 {
		t.Fatalf("actions=%d", len(got.Actions))
	}
}
