package main

import "path/filepath"

func resolveSnapshotPath(
	home string,
	arg string,
) string {

	dir := filepath.Join(
		home,
		".local",
		"share",
		"fedora-state",
		"snapshots",
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
