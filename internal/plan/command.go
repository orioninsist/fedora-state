package plan

import (
	"fmt"

	"fedora-state/internal/manifest"
)

func Summary(
	old manifest.Manifest,
	current manifest.Manifest,
) string {
	value := Build(old, current)

	return fmt.Sprintf(
		"actions=%d",
		len(value.Actions),
	)
}
