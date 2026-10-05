package plan

import (
	"testing"

	"fedora-state/internal/manifest"
)

func TestHasChangesFalseForSameManifest(t *testing.T) {
	value := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	}

	if HasChanges(value, value) {
		t.Fatal("did not expect changes")
	}
}
