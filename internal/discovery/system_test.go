package discovery

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOSRelease(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "os-release")

	content := `
# comment
NAME="Example Linux"
ID=example
VERSION='42 Stable'
VERSION_ID="42"
EMPTY=
`

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	got := osRelease(path)

	want := map[string]string{
		"NAME":       "Example Linux",
		"ID":         "example",
		"VERSION":    "42 Stable",
		"VERSION_ID": "42",
		"EMPTY":      "",
	}

	for key, expected := range want {
		if got[key] != expected {
			t.Errorf(
				"%s=%q, want %q",
				key,
				got[key],
				expected,
			)
		}
	}
}

func TestOSReleaseMissingFile(t *testing.T) {
	got := osRelease(
		filepath.Join(
			t.TempDir(),
			"missing-os-release",
		),
	)

	if len(got) != 0 {
		t.Fatalf("expected empty result, got %#v", got)
	}
}
