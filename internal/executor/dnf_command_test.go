package executor

import (
	"testing"

	"fedora-state/internal/plan"
)

func TestCommandDNFBackendSupportsActions(t *testing.T) {
	backend := NewCommandDNFBackend()

	for _, action := range []plan.Action{
		{
			Type:     "add",
			Identity: "rpm:bash:0:1.0-1:x86_64",
		},
		{
			Type:     "remove",
			Identity: "rpm:vim:0:1.0-1:x86_64",
		},
	} {
		if err := backend.Execute(action); err != nil {
			t.Fatal(err)
		}
	}
}
