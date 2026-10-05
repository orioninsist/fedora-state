package plan

import (
	"strings"
	"testing"

	"fedora-state/internal/manifest"
)

func TestSummaryShowsActionCount(t *testing.T) {
	got := Summary(
		manifest.Manifest{},
		manifest.Manifest{
			Entries: []manifest.Entry{
				{
					Identity: "rpm:bash:0:1.0-1:x86_64",
					Name:     "bash",
					Type:     "package",
				},
			},
		},
	)

	if !strings.Contains(got, "actions=1") {
		t.Fatalf("summary=%q", got)
	}
}
