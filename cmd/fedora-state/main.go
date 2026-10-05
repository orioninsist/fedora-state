package main

import (
	"os"
	"path/filepath"

	"fedora-state/internal/discovery"
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

	if err := report.WriteText(os.Stdout, system); err != nil {
		panic(err)
	}
}
