package main

import (
	"fmt"

	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/diff"
	"fedora-state/internal/persistence"
)

func runDiff(
	ctx commands.Context,
) error {

	if len(ctx.Args) == 0 {
		return fmt.Errorf("snapshot path required")
	}

	old, err := persistence.LoadSnapshot(
		ctx.Args[0],
	)

	if err != nil {
		return err
	}

	result := diff.Compare(
		old,
		ctx.System,
	)

	for _, object := range result.Added {
		fmt.Printf(
			"+ %s\n",
			object.Name,
		)
	}

	for _, object := range result.Removed {
		fmt.Printf(
			"- %s\n",
			object.Name,
		)
	}

	return nil
}
