package plan

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteText(t *testing.T) {
	var output bytes.Buffer

	err := WriteText(
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

	if !strings.Contains(
		output.String(),
		"add rpm:bash:0:1.0-1:x86_64",
	) {
		t.Fatalf("unexpected output: %q", output.String())
	}
}
