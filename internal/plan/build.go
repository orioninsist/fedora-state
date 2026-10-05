package plan

import "fedora-state/internal/manifest"

func Build(
	old manifest.Manifest,
	current manifest.Manifest,
) Plan {
	return Diff(old, current)
}
