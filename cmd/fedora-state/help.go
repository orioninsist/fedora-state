package main

import "fmt"

func printHelp() {
	fmt.Println(`Usage:

  fedora-state
  fedora-state --format=text
  fedora-state --format=json
  fedora-state --format=markdown
  fedora-state --format=manifest
  fedora-state --format=plan
  fedora-state --format=apply`)
}
