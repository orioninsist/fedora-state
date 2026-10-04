package hash

import (
	"testing"
	"time"

	"fedora-state/internal/model"
	"fedora-state/internal/state"
)

func TestStateIgnoresGeneratedAt(t *testing.T) {
	base := state.SystemState{
		OS: "fedora",
		Sources: []model.Source{
			{
				Name:      "rpm",
				Available: true,
				Collected: true,
			},
		},
		Packages: []model.Package{
			{
				Name:    "example",
				Version: "1.0",
				Manager: "rpm",
			},
		},
	}

	first := base
	first.GeneratedAt = time.Date(
		2026, 10, 4, 10, 0, 0, 0, time.UTC,
	)

	second := base
	second.GeneratedAt = time.Date(
		2026, 10, 5, 20, 0, 0, 0, time.UTC,
	)

	firstHash, err := State(first)
	if err != nil {
		t.Fatal(err)
	}

	secondHash, err := State(second)
	if err != nil {
		t.Fatal(err)
	}

	if firstHash != secondHash {
		t.Fatalf(
			"GeneratedAt changed state hash:\n%s\n%s",
			firstHash,
			secondHash,
		)
	}
}

func TestStateChangesWhenPackageChanges(t *testing.T) {
	first := state.SystemState{
		OS: "fedora",
		Packages: []model.Package{
			{
				Name:    "example",
				Version: "1.0",
				Manager: "rpm",
			},
		},
	}

	second := first
	second.Packages = []model.Package{
		{
			Name:    "example",
			Version: "2.0",
			Manager: "rpm",
		},
	}

	firstHash, err := State(first)
	if err != nil {
		t.Fatal(err)
	}

	secondHash, err := State(second)
	if err != nil {
		t.Fatal(err)
	}

	if firstHash == secondHash {
		t.Fatal("package change did not change state hash")
	}
}
