package report

import (
	"bytes"
	"strings"
	"testing"

	"fedora-state/internal/model"
)

func TestWriteText(t *testing.T) {
	var output bytes.Buffer

	err := WriteText(
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
		"os=linux",
		"architecture=amd64",
		"distribution=Fedora",
		"objects=1",
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
