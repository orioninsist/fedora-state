package executor

import "fmt"

func FormatResult(value Result) string {
	return fmt.Sprintf(
		"executed=%d skipped=%d",
		value.Executed,
		value.Skipped,
	)
}
