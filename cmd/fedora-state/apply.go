package main

import (
	"fedora-state/internal/executor"
	"fedora-state/internal/plan"
)

func applyPlan(value plan.Plan) (executor.Result, error) {
	return applyPlanWithEngine(
		newExecutorEngine(),
		value,
	)
}

func applyPlanWithEngine(
	engine executor.Engine,
	value plan.Plan,
) (executor.Result, error) {
	return engine.Apply(value)
}
