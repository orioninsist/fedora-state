package executor

import "fedora-state/internal/plan"

func DryRun(value plan.Plan) (Result, error) {
	return NewEngine(
		DryRunBackend{},
	).Apply(value)
}
