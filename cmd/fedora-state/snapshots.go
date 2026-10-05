package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"fedora-state/cmd/fedora-state/commands"
)

func runSnapshots(
	ctx commands.Context,
) error {

	dir := filepath.Join(
		ctx.Home,
		".local",
		"share",
		"fedora-state",
		"snapshots",
	)

	entries, err := os.ReadDir(dir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No snapshots found")
			return nil
		}

		return err
	}

	var names []string

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".json" {
			names = append(
				names,
				entry.Name(),
			)
		}
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
