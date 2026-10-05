package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestDNFBackendExecute(t *testing.T) {
	err := DNFBackend{}.Execute(plan.Action{
		Type:     "add",
		Identity: "rpm:bash:0:1.0-1:x86_64",
	})

	if err != nil {
		t.Fatal(err)
	}
}
