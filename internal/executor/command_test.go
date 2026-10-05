package executor

import (
	"strings"
	"testing"

	"fedora-state/internal/plan"
)

func TestSummaryShowsExecutionResult(t *testing.T) {
	got := Summary(plan.Plan{
		Actions: []plan.Action{
			{
				Type:     "add",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
		},
	})

	if !strings.Contains(got, "executed=1") {
		t.Fatalf("summary=%q", got)
	}
}
