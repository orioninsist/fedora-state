package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestDryRunBackendExecutesAction(t *testing.T) {
	err := DryRunBackend{}.Execute(plan.Action{
		Type:     "add",
		Identity: "rpm:bash:0:1.0-1:x86_64",
	})

	if err != nil {
		t.Fatal(err)
	}
}
