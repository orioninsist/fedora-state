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

	sort.Slice(
		entries,
		func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		},
	)

	fmt.Println("Snapshots:")

	for _, entry := range entries {

		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		info, err := entry.Info()

		if err != nil {
			return err
		}

		label := ""

		if entry.Name() == "latest.json" {
			label = " latest"
		}

		fmt.Printf(
			"%s\t%d bytes%s\n",
			entry.Name(),
			info.Size(),
			label,
		)
	}

	return nil
}
