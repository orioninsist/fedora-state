package uv

import (
	"os"
	"testing"
)

func TestParseReceipt(t *testing.T) {
	input := []byte(`
[tool]
requirements = [{ name = "gitfleet", git = "https://github.com/orioninsist/gitfleet.git" }]
entrypoints = [
    { name = "gitfleet", install-path = "/home/murat/.local/bin/gitfleet", from = "gitfleet" },
]
`)

	objects, err := parseReceipt(input)
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 1 {
		t.Fatalf("objects=%d want 1", len(objects))
	}

	object := objects[0]

	if object.Name != "gitfleet" {
		t.Fatalf("name=%q", object.Name)
	}

	if object.Identity != "uv:gitfleet:https://github.com/orioninsist/gitfleet.git" {
		t.Fatalf("identity=%q", object.Identity)
	}

	if len(object.Evidence) != 3 {
		t.Fatalf("evidence=%+v", object.Evidence)
	}

	foundExecutable := false
	for _, evidence := range object.Evidence {
		if evidence.Type == "executable_path" &&
			evidence.Value == "/home/murat/.local/bin/gitfleet" {
			foundExecutable = true
		}
	}
	if !foundExecutable {
		t.Fatalf("missing executable_path evidence: %+v", object.Evidence)
	}
}

func TestCollectorReadsReceipt(t *testing.T) {
	collector := NewCollector("/receipt")

	collector.read = func(path string) ([]byte, error) {
		if path != "/receipt" {
			t.Fatalf("path=%q", path)
		}

		return []byte(`
[tool]
requirements = [{ name = "gitfleet", git = "https://github.com/orioninsist/gitfleet.git" }]

entrypoints = [
	{ name = "gitfleet", install-path = "/home/murat/.local/bin/gitfleet", from = "gitfleet" },
]
`), nil
	}

	collection := collector.Collect()

	if collection.Err != nil {
		t.Fatal(collection.Err)
	}

	if len(collection.Objects) != 1 {
		t.Fatalf("objects=%d want 1", len(collection.Objects))
	}
}

func TestCollectorIgnoresMissingReceipt(t *testing.T) {
	collector := NewCollector("/missing")

	collector.read = func(string) ([]byte, error) {
		return nil, os.ErrNotExist
	}

	collection := collector.Collect()

	if collection.Err != nil {
		t.Fatal(collection.Err)
	}

	if len(collection.Objects) != 0 {
		t.Fatalf("objects=%d want 0", len(collection.Objects))
	}
}
