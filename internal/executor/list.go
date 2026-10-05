package executor

import (
	"strings"

	"fedora-state/internal/plan"
)

func DescribePlan(value plan.Plan) string {
	var lines []string

	for _, action := range value.Actions {
		lines = append(lines, Describe(action))
	}

	return strings.Join(lines, "\n")
}
