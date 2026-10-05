package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestExecuteWithBackend(t *testing.T) {
	backend := &RecorderBackend{}

	result, err := ExecuteWith(
		plan.Plan{
			Actions: []plan.Action{
				{
					Type:     "add",
					Identity: "rpm:bash:0:1.0-1:x86_64",
				},
			},
		},
		backend,
	)

	if err != nil {
		t.Fatal(err)
	}

	if result.Executed != 1 {
		t.Fatalf("executed=%d", result.Executed)
	}

	if len(backend.Actions) != 1 {
		t.Fatalf("actions=%d", len(backend.Actions))
	}
}
