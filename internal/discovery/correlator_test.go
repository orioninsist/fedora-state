package discovery

import (
	"testing"

	"fedora-state/internal/model"
)

func TestCorrelateExecutablePackageByPath(t *testing.T) {
	objects := []model.Object{
		{
			Name:     "tool",
			Type:     "executable",
			Location: "/home/test/.local/bin/tool",
			Evidence: []model.Evidence{
				{
					Type: "executable_path",
					Path: "/home/test/.local/bin/tool",
				},
			},
		},
		{
			Name:     "tool-package",
			Type:     "package",
			Identity: "example:tool-package:1",
			Evidence: []model.Evidence{
				{
					Type:  "executable_path",
					Value: "/home/test/.local/bin/tool",
				},
			},
		},
	}

	result := Correlate(objects)

	if len(result) != 2 {
		t.Fatalf("objects=%d want 2", len(result))
	}

	found := false
	for _, evidence := range result[0].Evidence {
		if evidence.Type == "package_identity" &&
			evidence.Value == "example:tool-package:1" {
			found = true
		}
	}

	if !found {
		t.Fatalf("missing package correlation: %+v", result[0].Evidence)
	}
}

func TestCorrelateDoesNotGuessByName(t *testing.T) {
	objects := []model.Object{
		{
			Name:     "tool",
			Type:     "executable",
			Location: "/usr/bin/tool",
		},
		{
			Name:     "tool",
			Type:     "package",
			Identity: "example:tool:1",
		},
	}

	result := Correlate(objects)

	for _, evidence := range result[0].Evidence {
		if evidence.Type == "package_identity" {
			t.Fatalf("guessed package correlation: %+v", evidence)
		}
	}
}

func TestCorrelateDoesNotMutateInput(t *testing.T) {
	objects := []model.Object{
		{
			Name:     "tool",
			Type:     "executable",
			Location: "/bin/tool",
		},
		{
			Name:     "package",
			Type:     "package",
			Identity: "example:package:1",
			Evidence: []model.Evidence{
				{
					Type:  "executable_path",
					Value: "/bin/tool",
				},
			},
		},
	}

	_ = Correlate(objects)

	if len(objects[0].Evidence) != 0 {
		t.Fatalf("input mutated: %+v", objects[0].Evidence)
	}
}
