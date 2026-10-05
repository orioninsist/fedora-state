package executor

import (
	"fmt"

	"fedora-state/internal/plan"
)

type RecorderBackend struct {
	Actions []plan.Action
}

func (b *RecorderBackend) Execute(action plan.Action) error {
	b.Actions = append(b.Actions, action)
	return nil
}

func ExecuteWith(
	value plan.Plan,
	backend Backend,
) (Result, error) {
	result := Result{}

	for _, action := range value.Actions {
		if err := backend.Execute(action); err != nil {
			return result, fmt.Errorf(
				"execute %s: %w",
				action.Identity,
				err,
			)
		}

		result.Executed++
	}

	return result, nil
}
