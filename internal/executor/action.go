package executor

import (
	"fmt"

	"fedora-state/internal/plan"
)

func Describe(value plan.Action) string {
	return fmt.Sprintf(
		"%s %s",
		value.Type,
		value.Identity,
	)
}
