package main

import (
	"testing"

	"fedora-state/internal/executor"
	"fedora-state/internal/manifest"
	"fedora-state/internal/plan"
)

func TestApplyFlowProducesStableState(t *testing.T) {
	old := manifest.Manifest{}

	current := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	}

	value := plan.Build(old, current)

	engine := executor.NewEngine(&executor.RecorderBackend{})

	result, err := engine.Apply(value)
	if err != nil {
		t.Fatal(err)
	}

	if result.Executed != 1 {
		t.Fatalf("executed=%d", result.Executed)
	}
}
