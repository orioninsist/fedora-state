package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {

	path := t.TempDir() + "/config.yaml"

	err := os.WriteFile(
		path,
		[]byte(`
snapshot:
  directory: /tmp/snapshots

restore:
  require_confirmation: true

providers:
  rpm: true
  cargo: false
`),
		0644,
	)

	if err != nil {
		t.Fatal(err)
	}

	value, err := Load(path)

	if err != nil {
		t.Fatal(err)
	}

	if value.Snapshot.Directory != "/tmp/snapshots" {
		t.Fatal("missing snapshot directory")
	}

	if !value.Restore.RequireConfirmation {
		t.Fatal("missing restore option")
	}

	if value.Providers.Cargo {
		t.Fatal("expected cargo disabled")
	}
}
