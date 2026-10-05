package main

import (
	"fedora-state/internal/executor"
	"fedora-state/internal/executor/providers"
)

func newExecutorEngine() executor.Engine {

	registry := providers.NewDefaultRegistry()

	backend := executor.NewRouterBackend(
		registry,
	)

	return executor.NewEngine(
		backend,
	)
}
