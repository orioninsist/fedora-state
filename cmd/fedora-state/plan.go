package main

import (
	"os"

	"fedora-state/internal/manifest"
	"fedora-state/internal/plan"
)

func writePlan(
	old manifest.Manifest,
	current manifest.Manifest,
) error {
	value := plan.Build(old, current)

	return plan.WriteJSON(
		os.Stdout,
		value,
	)
}
