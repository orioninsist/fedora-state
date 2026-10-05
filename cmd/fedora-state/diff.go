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
		resolveSnapshotPath(ctx.Home, ctx.Args[0]),
	)

	if err != nil {
		return err
	}

	result := diff.Compare(
		old,
		ctx.System,
	)

	if len(result.Added) == 0 &&
		len(result.Removed) == 0 {
		return nil
	}

	if len(result.Added) > 0 {
		fmt.Println("ADDED")
		fmt.Println()

		for _, object := range result.Added {
			fmt.Printf(
				"+ %s\n",
				object.Type,
			)

			fmt.Printf(
				"  name: %s\n",
				object.Name,
			)

			if object.Version != "" {
				fmt.Printf(
					"  version: %s\n",
					object.Version,
				)
			}

			fmt.Println()
		}
	}

	if len(result.Removed) > 0 {
		fmt.Println("REMOVED")
		fmt.Println()

		for _, object := range result.Removed {
			fmt.Printf(
				"- %s\n",
				object.Type,
			)

			fmt.Printf(
				"  name: %s\n",
				object.Name,
			)

			if object.Version != "" {
				fmt.Printf(
					"  version: %s\n",
					object.Version,
				)
			}

			fmt.Println()
		}
	}

	return nil
}
