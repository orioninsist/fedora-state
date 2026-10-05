package main

import (
	"fmt"

	"fedora-state/internal/discovery"
)

func main() {
	engine := discovery.NewWithCollectors(
		discovery.BinaryCollector{},
		discovery.NewRPMCollector(),
	)

	system := engine.Analyze()

	resolver := discovery.NewRPMOwnershipResolver()
	system.Objects = resolver.Resolve(system.Objects)

	fmt.Printf(
		"os=%s architecture=%s distribution=%s objects=%d\n",
		system.OS,
		system.Architecture,
		system.Metadata.Distribution,
		len(system.Objects),
	)
}
