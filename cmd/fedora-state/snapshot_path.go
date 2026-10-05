package main

import (
	"path/filepath"

	"fedora-state/internal/config"
)

func resolveSnapshotPath(
	home string,
	cfg config.Config,
	arg string,
) string {

	dir := snapshotDirectory(
		home,
		cfg,
	)

	if arg == "latest" {
		return filepath.Join(
			dir,
			"latest.json",
		)
	}

	if filepath.IsAbs(arg) {
		return arg
	}

	return filepath.Join(
		dir,
		arg,
	)
}
