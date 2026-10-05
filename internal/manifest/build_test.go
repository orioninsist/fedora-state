package manifest

import (
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestFromSystemIncludesPackagesOnly(t *testing.T) {
	got := FromSystem(model.System{
		Objects: []model.Object{
			{
				Name:     "bash",
				Type:     "package",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
			{
				Name: "bash",
				Type: "executable",
			},
		},
	})

	want := Manifest{
		Entries: []Entry{
			{
				Name:     "bash",
				Type:     "package",
				Identity: "rpm:bash:0:1.0-1:x86_64",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%#v want=%#v", got, want)
	}
}
