package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestRunnerBackendRejectsUnknownAction(t *testing.T) {
	err := NewRunnerBackend().Execute(plan.Action{
		Type:     "unknown",
		Identity: "test",
	})

	if err == nil {
		t.Fatal("expected error")
	}
}
