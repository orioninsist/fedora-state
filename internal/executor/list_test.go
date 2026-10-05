package executor

import (
	"strings"
	"testing"

	"fedora-state/internal/plan"
)

func TestDescribePlan(t *testing.T) {
	got := DescribePlan(plan.Plan{
		Actions: []plan.Action{
			{
				Type:     "add",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
		},
	})

	if !strings.Contains(
		got,
		"add rpm:bash:0:1.0-1:x86_64",
	) {
		t.Fatalf("plan=%q", got)
	}
}
