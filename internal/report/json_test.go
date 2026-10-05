package report

import (
	"bytes"
	"strings"
	"testing"

	"fedora-state/internal/model"
)

func TestWriteJSON(t *testing.T) {
	var output bytes.Buffer

	err := WriteJSON(
		&output,
		model.System{
			OS: "linux",
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
		`"os": "linux"`,
		`"name": "bash"`,
		`"type": "package"`,
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
