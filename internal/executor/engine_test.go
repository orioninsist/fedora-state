package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestEngineApply(t *testing.T) {
	engine := NewEngine(&RecorderBackend{})

	result, err := engine.Apply(plan.Plan{
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
