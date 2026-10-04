package discovery

import "fedora-state/internal/model"

type Engine struct {
	collectors []Collector
}

func NewWithCollectors(collectors ...Collector) Engine {
	return Engine{
		collectors: append([]Collector(nil), collectors...),
	}
}

func (e Engine) Analyze() model.System {
	system := System()

	for _, collector := range e.collectors {
		if collector == nil {
			continue
		}

		system.Objects = append(
			system.Objects,
			collector.Collect()...,
		)
	}

	return system
}
