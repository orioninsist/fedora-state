package providers

import (
	"fedora-state/internal/executor/providers/cargo"
	"fedora-state/internal/executor/providers/rpm"
	"fedora-state/internal/executor/providers/uv"
)

func NewDefaultRegistry() Registry {

	registry := NewRegistry()

	registry.Register(
		"rpm",
		rpm.New(),
	)

	registry.Register(
		"cargo",
		cargo.New(),
	)

	registry.Register(
		"uv",
		uv.New(),
	)

	return registry
}
