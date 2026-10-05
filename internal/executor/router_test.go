package executor

import (
	"testing"

	"fedora-state/internal/executor/providers"
	"fedora-state/internal/plan"
)

type testProvider struct {
	added bool
}

func (p *testProvider) Add(string) error {
	p.added = true
	return nil
}

func (p *testProvider) Remove(string) error {
	return nil
}


func TestRouterBackendRoutesAction(t *testing.T) {

	provider := &testProvider{}

	registry := providers.NewRegistry()

	registry.Register(
		"rpm",
		provider,
	)

	backend := NewRouterBackend(
		registry,
	)

	err := backend.Execute(
		plan.Action{
			Type: "add",
			Identity: "rpm:bash",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !provider.added {
		t.Fatal("provider not called")
	}
}
