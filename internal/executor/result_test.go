package executor

import (
	"strings"
	"testing"
)

func TestFormatResult(t *testing.T) {
	got := FormatResult(Result{
		Executed: 2,
		Skipped:  1,
	})

	if !strings.Contains(got, "executed=2") {
		t.Fatalf("result=%q", got)
	}
}
