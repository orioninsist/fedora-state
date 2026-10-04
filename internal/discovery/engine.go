package discovery

import "fedora-state/internal/model"

type Engine struct {
	collectors []Collector
}

func NewWithCollectors(collectors ...Collector) Engine {
	return Engine{
		collectors: collectors,
	}
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
