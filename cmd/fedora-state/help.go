package main

import "fmt"

func printHelp() {
	fmt.Println(`Usage:

  fedora-state

Commands:

  snapshot
      Create a system snapshot

  snapshots
      List available snapshots

  diff <snapshot>
      Compare current system with a snapshot

  restore <snapshot>
      Show restore plan

  restore <snapshot> --apply
      Apply restore plan

  manifest
      Print current package manifest

  plan
      Show current state changes

  apply
      Apply current plan

Output:

  fedora-state --format=text
  fedora-state --format=json
  fedora-state --format=markdown`)
}
