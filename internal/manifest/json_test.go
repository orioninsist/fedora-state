package manifest

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	var output bytes.Buffer

	err := WriteJSON(&output, Manifest{
		Entries: []Entry{
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	text := output.String()

	for _, expected := range []string{
		`"identity": "rpm:bash:0:1.0-1:x86_64"`,
		`"name": "bash"`,
		`"type": "package"`,
	} {
		if !strings.Contains(text, expected) {
			t.Fatalf("missing %q in %q", expected, text)
		}
	}
}
