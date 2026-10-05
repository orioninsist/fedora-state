package main

import (
	"testing"

	"fedora-state/internal/manifest"
	"fedora-state/internal/model"
	"fedora-state/internal/plan"
)

func TestRestoreBuildsPlanFromSnapshot(t *testing.T) {

	snapshot := model.System{
		Objects: []model.Object{
			{
				Name:     "neovim",
				Type:     "package",
				Identity: "rpm:neovim:0:1:x86_64",
			},
		},
	}

	target := manifest.FromSnapshot(snapshot)

	current := manifest.Manifest{}

	value := plan.Build(
		current,
		target,
	)

	if len(value.Actions) != 1 {
		t.Fatalf(
			"actions=%d",
			len(value.Actions),
		)
	}

	if value.Actions[0].Type != "add" {
		t.Fatalf(
			"type=%s",
			value.Actions[0].Type,
		)
	}
}
