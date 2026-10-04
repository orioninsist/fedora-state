package discovery

import (
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

type testCollector struct {
	objects []model.Object
}

func (collector testCollector) Collect() []model.Object {
	return collector.objects
}

func TestEngineCollectsObjectsWithoutKnowingTheirType(t *testing.T) {
	want := []model.Object{
		{
			Name:     "example",
			Type:     "test-object",
			Location: "/example",
			Evidence: []model.Evidence{
				{
					Type:  "test",
					Value: "evidence",
				},
			},
		},
	}

	system := NewWithCollectors(
		testCollector{
			objects: want,
		},
	).Analyze()

	if !reflect.DeepEqual(system.Objects, want) {
		t.Fatalf(
			"objects differ\n got: %#v\nwant: %#v",
			system.Objects,
			want,
		)
	}
}

func TestEngineCombinesCollectors(t *testing.T) {
	first := model.Object{
		Name: "first",
		Type: "one",
	}

	second := model.Object{
		Name: "second",
		Type: "two",
	}

	system := NewWithCollectors(
		testCollector{
			objects: []model.Object{first},
		},
		testCollector{
			objects: []model.Object{second},
		},
	).Analyze()

	want := []model.Object{
		first,
		second,
	}

	if !reflect.DeepEqual(system.Objects, want) {
		t.Fatalf(
			"objects differ\n got: %#v\nwant: %#v",
			system.Objects,
			want,
		)
	}
}

func TestEngineCanRunWithNoCollectors(t *testing.T) {
	system := NewWithCollectors().Analyze()

	if len(system.Objects) != 0 {
		t.Fatalf(
			"got unexpected objects: %#v",
			system.Objects,
		)
	}
}

func TestEngineRegistersCollectorsAtRuntime(t *testing.T) {
	engine := NewWithCollectors()

	first := model.Object{
		Name: "first",
		Type: "runtime",
	}

	second := model.Object{
		Name: "second",
		Type: "runtime",
	}

	engine.Register(
		testCollector{
			objects: []model.Object{first},
		},
	)

	engine.Register(
		testCollector{
			objects: []model.Object{second},
		},
	)

	system := engine.Analyze()

	want := []model.Object{
		first,
		second,
	}

	if !reflect.DeepEqual(system.Objects, want) {
		t.Fatalf(
			"objects differ\n got: %#v\nwant: %#v",
			system.Objects,
			want,
		)
	}
}

func TestEngineIgnoresNilCollector(t *testing.T) {
	engine := NewWithCollectors()

	engine.Register(nil)

	system := engine.Analyze()

	if len(system.Objects) != 0 {
		t.Fatalf(
			"got unexpected objects: %#v",
			system.Objects,
		)
	}
}
