package persistence

import (
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestSaveAndLoadSnapshot(t *testing.T) {

	path := t.TempDir() + "/snapshot.json"

	original := model.System{
		OS: "Fedora",
		Architecture: "x86_64",
		Objects: []model.Object{
			{
				Name: "neovim",
				Type: "package",
			},
		},
	}

	if err := SaveSnapshot(path, original); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadSnapshot(path)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(original, loaded) {
		t.Fatalf(
			"snapshot mismatch\n got=%#v\nwant=%#v",
			loaded,
			original,
		)
	}
}
