package persistence

import "fedora-state/internal/manifest"

func ManifestChanged(
	old manifest.Manifest,
	current manifest.Manifest,
) bool {
	return len(old.Entries) != len(current.Entries)
}
