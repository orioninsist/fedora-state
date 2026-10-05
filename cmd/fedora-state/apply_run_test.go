package main

import (
	"testing"

	"fedora-state/internal/manifest"
)

func TestRunApply(t *testing.T) {
	err := runApply(
		manifest.Manifest{},
		manifest.Manifest{},
	)

	if err != nil {
		t.Fatal(err)
	}
}
