package main

import (
	"fmt"

	"fedora-state/internal/discovery"
	"fedora-state/internal/providers/rpm"
)

func main() {
	engine := discovery.NewWithCollectors(
		discovery.BinaryCollector{},
		rpm.NewRPMCollector(),
	)

	system := engine.Analyze()

	resolver := rpm.NewRPMOwnershipResolver()
	system.Objects = resolver.Resolve(system.Objects)

	provenance := rpm.NewDNFProvenanceResolver()
	system.Objects = provenance.Resolve(system.Objects)

	fmt.Printf(
		"os=%s architecture=%s distribution=%s objects=%d\n",
		system.OS,
		system.Architecture,
		system.Metadata.Distribution,
		len(system.Objects),
	)
}
