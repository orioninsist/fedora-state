package discovery

import (
	"os/exec"

	"fedora-state/internal/model"
)

func Sources() []model.Source {

	checks := []string{
		"rpm",
		"flatpak",
		"cargo",
		"npm",
	}

	result := []model.Source{}

	for _, name := range checks {

		_, err := exec.LookPath(name)

		result = append(
			result,
			model.Source{
				Name:      name,
				Available: err == nil,
				Collected: name == "rpm" && err == nil,
			},
		)
	}

	return result
}
