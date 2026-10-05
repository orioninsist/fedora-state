package discovery

import "fedora-state/internal/model"

func Correlate(objects []model.Object) []model.Object {
	packagesByPath := make(map[string][]string)

	for _, object := range objects {
		if object.Type != "package" || object.Identity == "" {
			continue
		}

		for _, evidence := range object.Evidence {
			if evidence.Type != "executable_path" {
				continue
			}

			path := evidence.Path
			if path == "" {
				path = evidence.Value
			}
			if path == "" {
				continue
			}

			packagesByPath[path] = append(
				packagesByPath[path],
				object.Identity,
			)
		}
	}

	result := make([]model.Object, len(objects))

	for index, object := range objects {
		result[index] = object
		result[index].Evidence = append(
			[]model.Evidence(nil),
			object.Evidence...,
		)

		if object.Type != "executable" || object.Location == "" {
			continue
		}

		for _, identity := range packagesByPath[object.Location] {
			result[index].Evidence = append(
				result[index].Evidence,
				model.Evidence{
					Type:  "package_identity",
					Value: identity,
				},
			)
		}
	}

	return result
}
