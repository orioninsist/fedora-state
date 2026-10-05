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

		collection := collector.Collect()
		system.Objects = append(system.Objects, collection.Objects...)
		if collection.Err != nil {
			system.Diagnostics = append(system.Diagnostics, model.Diagnostic{Message: collection.Err.Error()})
		}
	}

	system.Objects = Correlate(system.Objects)

	return system
}
