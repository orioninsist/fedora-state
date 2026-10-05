package hash

import (
	"testing"

	"fedora-state/internal/model"
)

func TestStateChangesWhenObjectChanges(t *testing.T) {
	first := model.System{
		OS: "linux",
		Objects: []model.Object{
			{
				Name:     "example",
				Type:     "package",
				Identity: "rpm:example:0:1.0-1:x86_64",
				Version:  "1.0",
			},
		},
	}

	second := first
	second.Objects = append([]model.Object(nil), first.Objects...)
	second.Objects[0].Version = "2.0"

	firstHash, err := State(first)
	if err != nil {
		t.Fatal(err)
	}

	secondHash, err := State(second)
	if err != nil {
		t.Fatal(err)
	}

	if firstHash == secondHash {
		t.Fatal("object change did not change state hash")
	}
}

func TestStateIgnoresObjectAndEvidenceOrder(t *testing.T) {
	first := model.System{
		OS:           "linux",
		Architecture: "amd64",
		Objects: []model.Object{
			{
				Name:     "beta",
				Type:     "package",
				Identity: "cargo:beta:1.0:registry",
				Evidence: []model.Evidence{
					{Type: "source", Value: "registry"},
					{Type: "binary", Value: "beta"},
				},
			},
			{
				Name:     "alpha",
				Type:     "package",
				Identity: "rpm:alpha:0:1.0-1:x86_64",
			},
		},
		Diagnostics: []model.Diagnostic{
			{Message: "second"},
			{Message: "first"},
		},
	}

	second := first
	second.Objects = []model.Object{
		first.Objects[1],
		first.Objects[0],
	}
	second.Objects[1].Evidence = []model.Evidence{
		first.Objects[0].Evidence[1],
		first.Objects[0].Evidence[0],
	}
	second.Diagnostics = []model.Diagnostic{
		first.Diagnostics[1],
		first.Diagnostics[0],
	}

	firstHash, err := State(first)
	if err != nil {
		t.Fatal(err)
	}

	secondHash, err := State(second)
	if err != nil {
		t.Fatal(err)
	}

	if firstHash != secondHash {
		t.Fatalf("ordering changed state hash:\n%s\n%s", firstHash, secondHash)
	}
}
