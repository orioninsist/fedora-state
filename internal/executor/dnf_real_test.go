package executor

import (
	"os/exec"
	"testing"

	"fedora-state/internal/plan"
)

func TestRealDNFBackendBuildsCommands(t *testing.T) {
	backend := RealDNFBackend{
		command: func(name string, args ...string) *exec.Cmd {
			if name != "dnf" {
				t.Fatalf("command=%s", name)
			}

			return exec.Command(
				"true",
			)
		},
	}

	err := backend.Execute(plan.Action{
		Type:     "add",
		Identity: "bash",
	})

	if err != nil {
		t.Fatal(err)
	}
}
