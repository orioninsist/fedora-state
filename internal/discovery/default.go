package discovery

func New() Engine {
	return NewWithCollectors(
		defaultCollectors()...,
	)
}

func defaultCollectors() []Collector {
	return []Collector{
		BinaryCollector{},
	}
}
