package main

import (
	"path/filepath"

	"fedora-state/internal/config"
)

func snapshotDirectory(
	home string,
	cfg config.Config,
) string {

	if cfg.Snapshot.Directory != "" {
		return cfg.Snapshot.Directory
	}

	return filepath.Join(
		home,
		".local",
		"share",
		"fedora-state",
		"snapshots",
	)
}
