package discovery

import "fedora-state/internal/model"

func DiscoverObjects(binaries []model.Binary) []model.Object {
	result := make([]model.Object, 0, len(binaries))

	for _, binary := range binaries {
		evidence := []model.Evidence{
			{
				Type: "executable_path",
				Path: binary.Path,
			},
		}

		if binary.RealPath != "" && binary.RealPath != binary.Path {
			evidence = append(
				evidence,
				model.Evidence{
					Type: "resolved_path",
					Path: binary.RealPath,
				},
			)
		}

		if binary.Owner != "" {
			evidence = append(
				evidence,
				model.Evidence{
					Type:  "owner",
					Value: binary.Owner,
				},
			)
		}

		result = append(
			result,
			model.Object{
				Name:     binary.Name,
				Type:     "executable",
				Location: binary.Path,
				Evidence: evidence,
			},
		)
	}

	return result
}
