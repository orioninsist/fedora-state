package persistence

import (
	"reflect"
	"testing"

	"fedora-state/internal/manifest"
)

func TestSaveAndLoadManifest(t *testing.T) {
	path := t.TempDir() + "/manifest.json"

	original := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	}

	if err := SaveManifest(path, original); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadManifest(path)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(original, loaded) {
		t.Fatalf(
			"manifest mismatch\n got=%#v\nwant=%#v",
			loaded,
			original,
		)
	}
}
