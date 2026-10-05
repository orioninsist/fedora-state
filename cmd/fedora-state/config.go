package main

import (
	"os"
	"path/filepath"

	"fedora-state/internal/config"
)

func defaultConfigPath(home string) string {
	return filepath.Join(
		home,
		".config",
		"fedora-state",
		"config.yaml",
	)
}

func loadRuntimeConfig(
	home string,
) config.Config {

	path := defaultConfigPath(home)

	if _, err := os.Stat(path); err != nil {
		return config.Config{}
	}

	value, err := config.Load(path)

	if err != nil {
		return config.Config{}
	}

	return value
}
