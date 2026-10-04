package discovery

import "fedora-state/internal/model"

type BinaryCollector struct{}

func (BinaryCollector) Collect() []model.Object {
	return DiscoverObjects(DiscoverBinaries())
}
