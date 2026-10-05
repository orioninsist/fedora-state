package plan

import "fedora-state/internal/manifest"

func HasChanges(
	old manifest.Manifest,
	current manifest.Manifest,
) bool {
	return len(Build(old, current).Actions) > 0
}
