package plan

import "fedora-state/internal/manifest"

func Build(
	old manifest.Manifest,
	current manifest.Manifest,
) Plan {
	return Normalize(Diff(old, current))
}
