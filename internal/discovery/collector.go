package discovery

import "fedora-state/internal/model"

type Collector interface {
	Collect() []model.Object
}
