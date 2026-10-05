package main

import "fmt"

var (
	Version   = "dev"
	Commit    = "unknown"
	BuildDate = "unknown"
)

func printVersion() {
	fmt.Printf(
		"fedora-state %s\ncommit: %s\nbuilt: %s\n",
		Version,
		Commit,
		BuildDate,
	)
}
