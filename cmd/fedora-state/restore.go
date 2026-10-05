package main

import (
	"os"

	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/manifest"
	"fedora-state/internal/persistence"
	"fedora-state/internal/plan"
)

func runRestore(
	ctx commands.Context,
) error {

	if len(ctx.Args) == 0 {
		return os.ErrInvalid
	}

	snapshot, err := persistence.LoadSnapshot(
		resolveSnapshotPath(
			ctx.Home,
			ctx.Args[0],
		),
	)

	if err != nil {
		return err
	}

	target := manifest.FromSystem(snapshot)

	current := manifest.FromSystem(
		ctx.System,
	)

	value := plan.Build(
		current,
		target,
	)

	return plan.WriteJSON(
		os.Stdout,
		value,
	)
}
