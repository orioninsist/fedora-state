package executor

import (
	"fmt"

	"fedora-state/internal/plan"
)

func Summary(value plan.Plan) string {
	result, err := DryRun(value)

	if err != nil {
		return fmt.Sprintf(
			"error=%s",
			err,
		)
	}

	return fmt.Sprintf(
		"executed=%d skipped=%d",
		result.Executed,
		result.Skipped,
	)
}
