package main

import "testing"

func TestParseFormat(t *testing.T) {
	got := parseFormat([]string{
		"--format=plan",
	})

	if got != "plan" {
		t.Fatalf("format=%q", got)
	}
}
