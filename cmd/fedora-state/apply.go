package main

import (
	"fedora-state/internal/executor"
	"fedora-state/internal/plan"
)

func applyPlan(value plan.Plan) executor.Result {
	return executor.Run(value)
}
