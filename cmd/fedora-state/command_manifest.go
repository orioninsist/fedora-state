package main

import (
	"os"

	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/manifest"
)

func runManifest(ctx commands.Context) error {
	state := manifest.FromSystem(ctx.System)

	return manifest.WriteJSON(
		os.Stdout,
		state,
	)
}
