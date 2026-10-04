package discovery

import "fedora-state/internal/model"

type Collection struct {
	Objects []model.Object
	Err     error
}

// Collector discovers system objects from one evidence surface.
type Collector interface {
	Collect() Collection
}
