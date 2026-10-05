package executor

import (
	"fedora-state/internal/plan"
)

type DNFBackend struct {
}

func (DNFBackend) Execute(action plan.Action) error {
	return nil
}
