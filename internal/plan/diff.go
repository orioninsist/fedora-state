package plan

import "fedora-state/internal/manifest"

func Diff(
	old manifest.Manifest,
	newer manifest.Manifest,
) Plan {
	oldEntries := make(map[string]struct{})
	newEntries := make(map[string]struct{})

	for _, entry := range old.Entries {
		oldEntries[entry.Identity] = struct{}{}
	}

	for _, entry := range newer.Entries {
		newEntries[entry.Identity] = struct{}{}
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

	for _, entry := range old.Entries {
		if _, exists := newEntries[entry.Identity]; exists {
			continue
		}

		actions = append(actions, Action{
			Type:     "remove",
			Identity: entry.Identity,
		})
	}

	return Plan{
		Actions: actions,
	}
}
