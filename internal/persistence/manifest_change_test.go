package persistence

import (
	"testing"

	"fedora-state/internal/manifest"
)

func TestManifestChangedDetectsDifferentEntries(t *testing.T) {
	old := manifest.Manifest{}

	current := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	}

	if !ManifestChanged(old, current) {
		t.Fatal("expected manifest change")
	}
}
