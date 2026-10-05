package main

import (
	"fedora-state/cmd/fedora-state/commands"
)

func realHandlers() map[string]commands.Handler {
	return map[string]commands.Handler{
		commands.Help: func(commands.Context) error {
			printHelp()
			return nil
		},

		commands.Version: func(commands.Context) error {
			printVersion()
			return nil
		},

		commands.Manifest: func(ctx commands.Context) error {
			return nil
		},

		commands.Plan: func(ctx commands.Context) error {
			return nil
		},

		commands.Apply: func(ctx commands.Context) error {
			return nil
		},
	}
}
