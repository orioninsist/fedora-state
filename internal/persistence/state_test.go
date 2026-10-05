package persistence

import (
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestSaveAndLoadSystem(t *testing.T) {
	path := t.TempDir() + "/state.json"

	original := model.System{
		OS: "linux",
		Objects: []model.Object{
			{
				Name: "bash",
				Type: "package",
			},
		},
	}

	if err := SaveSystem(path, original); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadSystem(path)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(original, loaded) {
		t.Fatalf(
			"state mismatch\n got=%#v\nwant=%#v",
			loaded,
			original,
		)
	}
}
