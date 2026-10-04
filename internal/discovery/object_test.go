package discovery

import (
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestDiscoverObjectsPreservesBinaryEvidence(t *testing.T) {
	binaries := []model.Binary{
		{
			Name:     "tool",
			Path:     "/path/bin/tool",
			RealPath: "/real/bin/tool",
			Owner:    "root",
		},
	}

	got := DiscoverObjects(binaries)

	want := []model.Object{
		{
			Name:     "tool",
			Type:     "executable",
			Location: "/path/bin/tool",
			Evidence: []model.Evidence{
				{
					Type: "executable_path",
					Path: "/path/bin/tool",
				},
				{
					Type: "resolved_path",
					Path: "/real/bin/tool",
				},
				{
					Type:  "owner",
					Value: "root",
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(
			"objects differ\n got: %#v\nwant: %#v",
			got,
			want,
		)
	}
}

func TestDiscoverObjectsOmitsRedundantEvidence(t *testing.T) {
	binaries := []model.Binary{
		{
			Name:     "tool",
			Path:     "/bin/tool",
			RealPath: "/bin/tool",
		},
	}

	got := DiscoverObjects(binaries)

	if len(got) != 1 {
		t.Fatalf("got %d objects, want 1", len(got))
	}

	if len(got[0].Evidence) != 1 {
		t.Fatalf(
			"got %d evidence records, want 1: %#v",
			len(got[0].Evidence),
			got[0].Evidence,
		)
	}

	if got[0].Evidence[0].Type != "executable_path" {
		t.Fatalf(
			"evidence type=%q, want executable_path",
			got[0].Evidence[0].Type,
		)
	}
}

func TestDiscoverObjectsEmptyInput(t *testing.T) {
	got := DiscoverObjects(nil)

	if len(got) != 0 {
		t.Fatalf("got unexpected objects: %#v", got)
	}
}
