package main

import (
	"fmt"
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
			ctx.Config,
			ctx.Args[0],
		),
	)

	if err != nil {
		return err
	}

	target := manifest.FromSnapshot(snapshot)

	current := manifest.FromSystem(
		ctx.System,
	)

	value := plan.Build(
		current,
		target,
	)

	if !restoreApplyRequested(ctx.Args) {
		return plan.WriteJSON(
			os.Stdout,
			value,
		)
	}

	engine := newExecutorEngine()

	result, err := engine.Apply(value)

	if err != nil {
		return err
	}

	fmt.Printf(
		"executed=%d\n",
		result.Executed,
	)

	return nil
}

func restoreApplyRequested(
	args []string,
) bool {

	for _, arg := range args {
		if arg == "--apply" {
			return true
		}
	}

	return false
}
