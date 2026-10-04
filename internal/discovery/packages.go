package discovery

import (
	"fedora-state/internal/model"
)

func Packages(
	sources []model.Source,
) ([]model.Package, error) {

	result := []model.Package{}

	for _, source := range sources {

		if !source.Available {
			continue
		}

		switch source.Name {

		case "rpm":

			packages, err := RPMPackages()

			if err != nil {
				return result, err
			}

			result = append(
				result,
				packages...,
			)

		}
	}

	return result, nil
}
