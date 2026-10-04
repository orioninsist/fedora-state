package discovery

import "fedora-state/internal/model"

type Engine struct {
	collectors []Collector
}

func NewWithCollectors(collectors ...Collector) Engine {
	engine := Engine{}

	for _, collector := range collectors {
		engine.Register(collector)
	}

	return engine
}

func (e *Engine) Register(collector Collector) {
	if collector == nil {
		return
	}

	e.collectors = append(
		e.collectors,
		collector,
	)
}

func (e Engine) Analyze() model.System {
	system := System()

	for _, collector := range e.collectors {
		system.Objects = append(
			system.Objects,
			collector.Collect()...,
		)
	}

	return system
}
