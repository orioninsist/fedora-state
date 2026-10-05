package main

import (
	"fedora-state/internal/executor"
	"fedora-state/internal/plan"
)

func applyPlan(value plan.Plan) (executor.Result, error) {
	return newExecutorEngine().Apply(value)
}
