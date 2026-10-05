package cargo

import (
	"os"
	"testing"
)

func TestParseObjects(t *testing.T) {
	input := []byte(`{"installs":{"wl-screenrec 0.2.0 (registry+https://github.com/rust-lang/crates.io-index)":{"bins":["wl-screenrec"],"target":"x86_64-unknown-linux-gnu"}}}`)

	objects, err := parseObjects(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(objects) != 1 {
		t.Fatalf("objects=%d, want 1", len(objects))
	}

	object := objects[0]
	if object.Name != "wl-screenrec" || object.Version != "0.2.0" {
		t.Fatalf("object=%+v", object)
	}
	if object.Identity != "cargo:wl-screenrec:0.2.0:registry+https://github.com/rust-lang/crates.io-index" {
		t.Fatalf("identity=%q", object.Identity)
	}
	if len(object.Evidence) != 4 {
		t.Fatalf("evidence=%+v", object.Evidence)
	}

	foundExecutable := false
	for _, evidence := range object.Evidence {
		if evidence.Type == "executable_name" &&
			evidence.Value == "wl-screenrec" {
			foundExecutable = true
		}
	}
	if !foundExecutable {
		t.Fatalf("missing executable_name evidence: %+v", object.Evidence)
	}
}

func TestParseObjectsRejectsInvalidJSON(t *testing.T) {
	if _, err := parseObjects([]byte(`{`)); err == nil {
		t.Fatal("expected JSON error")
	}
}

func TestCollectorReadsMetadata(t *testing.T) {
	collector := NewCollector("/metadata")
	collector.read = func(path string) ([]byte, error) {
		if path != "/metadata" {
			t.Fatalf("path=%q", path)
		}
		return []byte(`{"installs":{"wl-screenrec 0.2.0 (registry+https://example.test/index)":{"bins":["wl-screenrec"]}}}`), nil
	}

	collection := collector.Collect()
	if collection.Err != nil {
		t.Fatal(collection.Err)
	}
	if len(collection.Objects) != 1 {
		t.Fatalf("objects=%d, want 1", len(collection.Objects))
	}
}

func TestCollectorIgnoresMissingMetadata(t *testing.T) {
	collector := NewCollector("/missing")
	collector.read = func(string) ([]byte, error) {
		return nil, os.ErrNotExist
	}

	collection := collector.Collect()
	if collection.Err != nil || len(collection.Objects) != 0 {
		t.Fatalf("collection=%+v", collection)
	}
}
