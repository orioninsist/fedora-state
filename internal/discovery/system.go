package discovery

import (
	"os"
	"os/exec"

	"fedora-state/internal/model"
)

func DiscoverSources() []string {

	sources := []string{}

	checks := map[string]string{
		"rpm":     "rpm",
		"flatpak": "flatpak",
		"cargo":   "cargo",
		"npm":     "npm",
	}

	for name, cmd := range checks {

		if _, err := exec.LookPath(cmd); err == nil {
			sources = append(
				sources,
				name,
			)
		}
	}

	return sources
}

func DetectOS() string {

	if _, err := os.Stat("/etc/fedora-release"); err == nil {
		return "fedora"
	}

	return "unknown"
}

func System() model.SystemState {

	return model.SystemState{
		OS:      DetectOS(),
		Sources: DiscoverSources(),
	}
}
