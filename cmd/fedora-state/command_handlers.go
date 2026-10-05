package main

import (
	"os"

	"fedora-state/cmd/fedora-state/commands"
	"fedora-state/internal/report"
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

		commands.Manifest: runManifest,

		commands.Plan: runPlan,

		commands.Apply: runApplyCommand,

		commands.Snapshot: runSnapshot,

		commands.Snapshots: runSnapshots,

		"text": func(ctx commands.Context) error {
			return report.WriteText(os.Stdout, ctx.System)
		},

		"json": func(ctx commands.Context) error {
			return report.WriteJSON(os.Stdout, ctx.System)
		},

		"markdown": func(ctx commands.Context) error {
			return report.WriteMarkdown(os.Stdout, ctx.System)
		},
	}
}
