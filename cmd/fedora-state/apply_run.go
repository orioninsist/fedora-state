package main

import (
	"fedora-state/internal/manifest"
	"fedora-state/internal/persistence"
	"fedora-state/internal/plan"
)

func runApply(
	old manifest.Manifest,
	current manifest.Manifest,
) error {
	value := plan.Build(old, current)

	engine := newExecutorEngine()

	_, err := engine.Apply(value)
	if err != nil {
		return err
	}

	return persistence.SaveManifest(
		manifestPath(),
		current,
	)
}
