package main

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestApplyPlan(t *testing.T) {
	result, err := applyPlan(plan.Plan{
		Actions: []plan.Action{
			{
				Type:     "add",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if result.Executed != 1 {
		t.Fatalf("executed=%d", result.Executed)
	}
}
