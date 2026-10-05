package manifest

import "fedora-state/internal/model"

func FromSnapshot(
	system model.System,
) Manifest {

	return FromSystem(system)
}
