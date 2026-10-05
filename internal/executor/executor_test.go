package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestApplyCountsSupportedActions(t *testing.T) {
	result := New().Apply(plan.Plan{
		Actions: []plan.Action{
			{
				Type:     "add",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
			{
				Type:     "remove",
				Identity: "rpm:vim:0:1.0-1:x86_64",
			},
		},
	})

	if result.Executed != 2 {
		t.Fatalf("executed=%d", result.Executed)
	}
}
