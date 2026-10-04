package discovery

import (
	"errors"
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

type testCollector struct {
	objects []model.Object
	err     error
}

func (collector testCollector) Collect() Collection {
	return Collection{Objects: collector.objects, Err: collector.err}
}

func TestEngineCollectsObjectsWithoutKnowingTheirType(t *testing.T) {
	want := []model.Object{
		{
			Name:     "example",
			Type:     "test-object",
			Identity: "example:stable-id",
			Location: "/example",
			Version:  "1.2.3",
			Source:   "test-source",
			Evidence: []model.Evidence{
				{
					Type:  "test",
					Value: "evidence",
				},
			},
		},
	}

	system := NewWithCollectors(
		testCollector{objects: want},
	).Analyze()

	if !reflect.DeepEqual(system.Objects, want) {
		t.Fatalf(
			"objects differ\n got: %#v\nwant: %#v",
			system.Objects,
			want,
		)
	}
}

func TestEngineCombinesCollectorsWithoutInterpretingObjects(t *testing.T) {
	first := model.Object{
		Name:     "first",
		Type:     "arbitrary-a",
		Identity: "a:first",
	}

	second := model.Object{
		Name:     "second",
		Type:     "arbitrary-b",
		Identity: "b:second",
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

func TestEnginePreservesCollectorEvidenceExactly(t *testing.T) {
	want := model.Object{
		Name:     "tool",
		Type:     "anything",
		Identity: "anything:tool",
		Version:  "9.8.7",
		Source:   "origin",
		Evidence: []model.Evidence{
			{
				Type:  "registry",
				Path:  "/var/lib/example/database",
				Value: "record",
			},
			{
				Type:  "receipt",
				Path:  "/example/receipt",
				Value: "source metadata",
			},
		},
	}

	system := NewWithCollectors(
		testCollector{
			objects: []model.Object{want},
		},
	).Analyze()

	if len(system.Objects) != 1 {
		t.Fatalf(
			"objects=%d want=1",
			len(system.Objects),
		)
	}

	if !reflect.DeepEqual(system.Objects[0], want) {
		t.Fatalf(
			"object changed by engine\n got: %#v\nwant: %#v",
			system.Objects[0],
			want,
		)
	}
}

func TestEngineIgnoresNilCollectors(t *testing.T) {
	want := model.Object{
		Name: "example",
		Type: "test",
	}

	system := NewWithCollectors(
		nil,
		testCollector{
			objects: []model.Object{want},
		},
		nil,
	).Analyze()

	if !reflect.DeepEqual(system.Objects, []model.Object{want}) {
		t.Fatalf(
			"objects differ\n got: %#v\nwant: %#v",
			system.Objects,
			[]model.Object{want},
		)
	}
}

func TestEnginePreservesCollectorErrorsAsDiagnostics(t *testing.T) {
	want := model.Object{Name: "survivor", Type: "test"}
	system := NewWithCollectors(testCollector{err: errors.New("collection failed")}, testCollector{objects: []model.Object{want}}).Analyze()
	if !reflect.DeepEqual(system.Objects, []model.Object{want}) {
		t.Fatalf("objects differ: %#v", system.Objects)
	}
	if len(system.Diagnostics) != 1 || system.Diagnostics[0].Message != "collection failed" {
		t.Fatalf("diagnostics differ: %#v", system.Diagnostics)
	}
}
