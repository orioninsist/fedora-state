package report

import (
	"bytes"
	"strings"
	"testing"

	"fedora-state/internal/model"
)

func TestWriteMarkdown(t *testing.T) {
	var output bytes.Buffer

	err := WriteMarkdown(
		&output,
		model.System{
			OS:           "linux",
			Architecture: "amd64",
			Metadata: model.SystemMetadata{
				Distribution: "Fedora",
			},
			Objects: []model.Object{
				{
					Name: "bash",
					Type: "package",
				},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	text := output.String()

	for _, expected := range []string{
		"# Fedora State Report",
		"OS: linux",
		"Distribution: Fedora",
		"`bash` (package)",
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf(
				"missing %q in %q",
				expected,
				text,
			)
		}
	}
}
