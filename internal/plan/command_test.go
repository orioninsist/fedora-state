package plan

import (
	"testing"

	"fedora-state/internal/manifest"
)

func TestBuildPlanKeepsRemoveActions(t *testing.T) {
	old := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:vim:0:1.0-1:x86_64",
				Name:     "vim",
				Type:     "package",
			},
		},
	}

	current := manifest.Manifest{
		Entries: []manifest.Entry{},
	}

	got := Build(old, current)

	if len(got.Actions) != 1 {
		t.Fatalf("actions=%d", len(got.Actions))
	}

	if got.Actions[0].Type != "remove" {
		t.Fatalf("type=%q", got.Actions[0].Type)
	}
}
