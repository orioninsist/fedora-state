package diff

import "fedora-state/internal/model"

type Result struct {
	Added   []model.Object
	Removed []model.Object
}

func Compare(
	old model.System,
	current model.System,
) Result {

	oldObjects := map[string]model.Object{}

	for _, object := range old.Objects {
		if object.Identity != "" {
			oldObjects[object.Identity] = object
		}
	}

	currentObjects := map[string]model.Object{}

	for _, object := range current.Objects {
		if object.Identity != "" {
			currentObjects[object.Identity] = object
		}
	}

	var result Result

	for id, object := range currentObjects {
		if _, ok := oldObjects[id]; !ok {
			result.Added = append(
				result.Added,
				object,
			)
		}
	}

	for id, object := range oldObjects {
		if _, ok := currentObjects[id]; !ok {
			result.Removed = append(
				result.Removed,
				object,
			)
		}
	}

	return result
}
