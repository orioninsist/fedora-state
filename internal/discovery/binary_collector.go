package discovery

type BinaryCollector struct{}

func (BinaryCollector) Collect() Collection {
	return Collection{Objects: DiscoverObjects(DiscoverBinaries())}
}
