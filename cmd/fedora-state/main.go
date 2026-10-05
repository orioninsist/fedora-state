package main

import (
	"os"
	"path/filepath"
	"strings"

	"fedora-state/internal/discovery"
	"fedora-state/internal/hash"
	"fedora-state/internal/persistence"
	"fedora-state/internal/providers/cargo"
	"fedora-state/internal/providers/rpm"
	"fedora-state/internal/providers/uv"

	"fedora-state/cmd/fedora-state/commands"
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
		format = commands.CommandFromArgs([]string{
			strings.TrimPrefix(os.Args[1], "--format="),
		})
	}

	if err := commands.Dispatch(
		format,
		commands.Context{
			Args:   os.Args[2:],
			Home:   home,
			System: system,
		},
		realHandlers(),
	); err != nil {
		panic(err)
	}

}
