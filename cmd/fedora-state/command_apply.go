package main

import (
	"os"
	"path/filepath"

	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/manifest"
	"fedora-state/internal/persistence"
	"fedora-state/internal/plan"
)

func runApplyCommand(ctx commands.Context) error {
	current := manifest.FromSystem(ctx.System)

	old, err := persistence.LoadManifest(
		filepath.Join(
			ctx.Home,
			".local",
			"state",
			"fedora-state-manifest.json",
		),
	)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	value := plan.Build(old, current)

	_ = applyPlan(value)

	return persistence.SaveManifest(
		filepath.Join(
			ctx.Home,
			".local",
			"state",
			"fedora-state-manifest.json",
		),
		current,
	)
}
