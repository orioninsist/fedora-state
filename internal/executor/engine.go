package executor

import "fedora-state/internal/plan"

type Engine struct {
	backend Backend
}

func NewEngine(
	backend Backend,
) Engine {
	return Engine{
		backend: backend,
	}
}

func (e Engine) Apply(
	value plan.Plan,
) (Result, error) {
	return ExecuteWith(
		value,
		e.backend,
	)
}
