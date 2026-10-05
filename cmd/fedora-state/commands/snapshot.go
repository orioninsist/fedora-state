package commands

import (
	"os"
	"path/filepath"

	"fedora-state/internal/persistence"
)

func WriteSnapshot(ctx Context) error {

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
