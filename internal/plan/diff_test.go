package plan

import (
	"reflect"
	"testing"

	"fedora-state/internal/manifest"
)

func TestDiffAddsMissingManifestEntry(t *testing.T) {
	old := manifest.Manifest{
		Entries: []manifest.Entry{},
	}

	newer := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	}

	got := Diff(old, newer)

	want := Plan{
		Actions: []Action{
			{
				Type:     "add",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%#v want=%#v", got, want)
	}
}
