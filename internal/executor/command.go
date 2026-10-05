package executor

import (
	"fmt"

	"fedora-state/internal/plan"
)

func Summary(value plan.Plan) string {
	result := New().Apply(value)

	return fmt.Sprintf(
		"executed=%d skipped=%d",
		result.Executed,
		result.Skipped,
	)
}
