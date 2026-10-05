package discovery

import "fedora-state/internal/model"

type Resolver interface {
	Resolve([]model.Object) []model.Object
}
