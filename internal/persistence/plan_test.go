package persistence

import (
	"reflect"
	"testing"

	"fedora-state/internal/plan"
)

func TestSaveAndLoadPlan(t *testing.T) {
	path := t.TempDir() + "/plan.json"

	original := plan.Plan{
		Actions: []plan.Action{
			{
				Type:     "add",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
		},
	}

	if err := SavePlan(path, original); err != nil {
		t.Fatal(err)
	}

	loaded, err := LoadPlan(path)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(original, loaded) {
		t.Fatalf(
			"plan mismatch\n got=%#v\nwant=%#v",
			loaded,
			original,
		)
	}
}
