package main

import (
	"os"
	"path/filepath"
	"strings"

	"fedora-state/internal/discovery"
	"fedora-state/internal/hash"
	"fedora-state/internal/manifest"
	"fedora-state/internal/persistence"
	"fedora-state/internal/plan"
	"fedora-state/internal/providers/cargo"
	"fedora-state/internal/providers/rpm"
	"fedora-state/internal/providers/uv"
	"fedora-state/internal/report"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	engine := discovery.NewWithCollectors(
		discovery.BinaryCollector{},
		rpm.NewRPMCollector(),
		cargo.NewCollector(filepath.Join(home, ".local", ".crates2.json")),
		uv.NewScanner(filepath.Join(home, ".local", "share", "uv", "tools")),
	)

	system := engine.Analyze()

	resolver := rpm.NewRPMOwnershipResolver()
	system.Objects = resolver.Resolve(system.Objects)

	provenance := rpm.NewDNFProvenanceResolver()
	system.Objects = provenance.Resolve(system.Objects)

	stateHash, err := hash.State(system)
	if err != nil {
		panic(err)
	}

	stateHashPath := filepath.Join(home, ".local", "state", "fedora-state.hash")

	oldHash, err := persistence.LoadHash(stateHashPath)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	if persistence.Changed(oldHash, stateHash) {
		if err := os.MkdirAll(filepath.Dir(stateHashPath), 0755); err != nil {
			panic(err)
		}

		if err := persistence.SaveHash(stateHashPath, stateHash); err != nil {
			panic(err)
		}
	}

	format := "text"
	if len(os.Args) > 1 {
		format = strings.TrimPrefix(os.Args[1], "--format=")
	}

	if format == "plan" {
		current := manifest.FromSystem(system)
		old, err := persistence.LoadManifest(
			filepath.Join(home, ".local", "state", "fedora-state-manifest.json"),
		)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}

		value := plan.Build(old, current)

		if err := plan.WriteJSON(os.Stdout, value); err != nil {
			panic(err)
		}
		return
	}

	switch format {
	case "json":
		if err := report.WriteJSON(os.Stdout, system); err != nil {
			panic(err)
		}
	case "manifest":
		state := manifest.FromSystem(system)
		if err := manifest.WriteJSON(os.Stdout, state); err != nil {
			panic(err)
		}
	default:
		if err := report.WriteText(os.Stdout, system); err != nil {
			panic(err)
		}
	}
}
