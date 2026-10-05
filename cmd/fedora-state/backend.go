package main

import "fedora-state/internal/executor"

func newExecutorEngine() executor.Engine {
	return executor.NewEngine(
		executor.NewRealDNFBackend(),
	)
}
