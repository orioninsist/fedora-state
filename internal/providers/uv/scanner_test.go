package uv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScannerCollectsReceipts(t *testing.T) {
	root := t.TempDir()

	dir := filepath.Join(root, "gitfleet")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}

	err := os.WriteFile(
		filepath.Join(dir, "uv-receipt.toml"),
		[]byte(`
[tool]
requirements = [{ name = "gitfleet", git = "https://github.com/orioninsist/gitfleet.git" }]

entrypoints = [
	{ name = "gitfleet", install-path = "/home/test/.local/bin/gitfleet", from = "gitfleet" },
]
`),
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}

	collection := NewScanner(root).Collect()

	if collection.Err != nil {
		t.Fatal(collection.Err)
	}

	if len(collection.Objects) != 1 {
		t.Fatalf("objects=%d want 1", len(collection.Objects))
	}
}

func TestScannerIgnoresMissingDirectory(t *testing.T) {
	collection := NewScanner("/missing/uv/tools").Collect()

	if collection.Err != nil {
		t.Fatal(collection.Err)
	}

	if len(collection.Objects) != 0 {
		t.Fatalf("objects=%d want 0", len(collection.Objects))
	}
}
