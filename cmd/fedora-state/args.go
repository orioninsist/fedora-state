package main

import "strings"

func parseFormat(args []string) string {
	if len(args) == 0 {
		return "text"
	}

	return strings.TrimPrefix(
		args[0],
		"--format=",
	)
}
