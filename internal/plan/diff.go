package plan

import "fedora-state/internal/manifest"

func Diff(
	old manifest.Manifest,
	newer manifest.Manifest,
) Plan {
	oldEntries := make(map[string]struct{})

	for _, entry := range old.Entries {
		oldEntries[entry.Identity] = struct{}{}
	}

	var actions []Action

	for _, entry := range newer.Entries {
		if _, exists := oldEntries[entry.Identity]; exists {
			continue
		}

		actions = append(actions, Action{
			Type:     "add",
			Identity: entry.Identity,
		})
	}

	return Plan{
		Actions: actions,
	}
}
