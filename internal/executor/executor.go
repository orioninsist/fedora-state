package executor

import "fedora-state/internal/plan"

type Executor struct{}

func New() Executor {
	return Executor{}
}

func (Executor) Apply(value plan.Plan) Result {
	result := Result{}

	for _, action := range value.Actions {
		switch action.Type {
		case "add", "remove":
			result.Executed++
		default:
			result.Skipped++
		}
	}

	return result
}
