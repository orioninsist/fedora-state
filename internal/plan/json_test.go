package plan

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	var output bytes.Buffer

	err := WriteJSON(
		&output,
		Plan{
			Actions: []Action{
				{
					Type:     "add",
					Identity: "rpm:bash:0:1.0-1:x86_64",
				},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	text := output.String()

	for _, expected := range []string{
		`"type": "add"`,
		`"identity": "rpm:bash:0:1.0-1:x86_64"`,
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf("missing %q in %q", expected, text)
		}
	}
}
