package discovery

import "fedora-state/internal/model"

type Engine struct{}

func New() Engine {
	return Engine{}
}

func (e Engine) Analyze() model.System {
	system := System()

	system.Binaries = DiscoverBinaries()
	system.Objects = DiscoverObjects(system.Binaries)

	return system
}
