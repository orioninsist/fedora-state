package main

import (
	"os"
	"path/filepath"
)

func manifestPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(
		home,
		".local",
		"state",
		"fedora-state-manifest.json",
	)
}
