package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/persistence"
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

	name := fmt.Sprintf(
		"%s.json",
		time.Now().Format("20060102-150405"),
	)

	path := filepath.Join(
		dir,
		name,
	)

	if err := persistence.SaveSnapshot(
		path,
		ctx.System,
	); err != nil {
		return err
	}

	latest := filepath.Join(
		dir,
		"latest.json",
	)

	return persistence.SaveSnapshot(
		latest,
		ctx.System,
	)
}
