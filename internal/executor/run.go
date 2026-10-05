package executor

import "fedora-state/internal/plan"

func Run(value plan.Plan) Result {
	return New().Apply(value)
}
