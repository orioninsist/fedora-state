package executor

import "fedora-state/internal/plan"

func Run(value plan.Plan) (Result, error) {
	return New().Apply(value), nil
}
