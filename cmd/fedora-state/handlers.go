package main

import (
	"os"

	"fedora-state/internal/discovery"
	"fedora-state/internal/manifest"
	"fedora-state/internal/report"

	"fedora-state/cmd/fedora-state/commands"
)

func commandHandlers(system discovery.System, home string) map[string]commands.Handler {
	return map[string]commands.Handler{
		commands.Help: func(commands.Context) error {
			printHelp()
			return nil
		},

		commands.Version: func(commands.Context) error {
			printVersion()
			return nil
		},

		commands.Manifest: func(commands.Context) error {
			state := manifest.FromSystem(system)

			return manifest.WriteJSON(
				os.Stdout,
				state,
			)
		},

		commands.Plan: func(commands.Context) error {
			return nil
		},

		commands.Apply: func(commands.Context) error {
			return nil
		},
	}
}
