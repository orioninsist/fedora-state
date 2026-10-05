package executor

import (
	"fedora-state/internal/executor/providers"
	"fedora-state/internal/plan"
)

type RouterBackend struct {
	registry providers.Registry
}

func NewRouterBackend(
	registry providers.Registry,
) RouterBackend {
	return RouterBackend{
		registry: registry,
	}
}

func (b RouterBackend) Execute(
	action plan.Action,
) error {

	identity := providers.Parse(
		action.Identity,
	)

	handler, err := b.registry.Lookup(
		identity.Provider,
	)

	if err != nil {
		return err
	}

	switch action.Type {

	case "add":
		return handler.Add(
			identity.Value,
		)

	case "remove":
		return handler.Remove(
			identity.Value,
		)

	default:
		return nil
	}
}
