package main

import (
	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/persistence"
	"os"
	"path/filepath"
)

func runSnapshot(
	ctx commands.Context,
) error {

	dir := filepath.Join(
		ctx.Home,
		".local",
		"share",
		"fedora-state",
		"snapshots",
	)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(
		dir,
		"machine.json",
	)

	return persistence.SaveSnapshot(
		path,
		ctx.System,
	)
}
