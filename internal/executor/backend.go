package executor

import "fedora-state/internal/plan"

type Backend interface {
	Execute(plan.Action) error
}

type DryRunBackend struct {
}

func (DryRunBackend) Execute(action plan.Action) error {
	return nil
}
