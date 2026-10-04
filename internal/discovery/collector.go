package discovery

import "fedora-state/internal/model"

// Collector discovers system objects from one evidence surface.
//
// The engine intentionally knows nothing about the mechanism behind a
// collector. Package databases, receipts, registries, filesystem metadata
// and other evidence surfaces all satisfy the same contract.
type Collector interface {
	Collect() []model.Object
}
