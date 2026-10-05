package discovery

import (
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

type staticCollector struct {
	objects []model.Object
}

func (collector staticCollector) Collect() Collection {
	return Collection{
		Objects: collector.objects,
	}
}

func TestAnalyzeCurrentSystem(t *testing.T) {
	system := NewWithCollectors(BinaryCollector{}).Analyze()

	t.Logf("os=%s", system.OS)
	t.Logf("architecture=%s", system.Architecture)
	t.Logf("distribution=%q", system.Metadata.Distribution)
	t.Logf("distribution_id=%q", system.Metadata.DistributionID)
	t.Logf("version=%q", system.Metadata.Version)
	t.Logf("version_id=%q", system.Metadata.VersionID)
	t.Logf("kernel=%q", system.Metadata.Kernel)
	t.Logf("hostname=%q", system.Metadata.Hostname)
	t.Logf("objects=%d", len(system.Objects))

	if system.OS == "" {
		t.Error("OS is empty")
	}

	if system.Architecture == "" {
		t.Error("architecture is empty")
	}

	if len(system.Objects) == 0 {
		t.Error("no objects discovered")
	}
}

func TestAnalyzeIsDeterministic(t *testing.T) {
	first := NewWithCollectors(BinaryCollector{}).Analyze()
	second := NewWithCollectors(BinaryCollector{}).Analyze()

	if !reflect.DeepEqual(first, second) {
		t.Error("consecutive discovery results differ")
	}
}

func TestAnalyzeCorrelatesExecutableWithPackage(t *testing.T) {
	collectors := NewWithCollectors(
		staticCollector{
			objects: []model.Object{
				{
					Name:     "tool",
					Type:     "executable",
					Location: "/usr/bin/tool",
				},
				{
					Name:     "tool",
					Type:     "package",
					Identity: "rpm:tool:0:1.0-1:x86_64",
					Evidence: []model.Evidence{
						{
							Type:  "executable_path",
							Value: "/usr/bin/tool",
						},
					},
				},
			},
		},
	)

	system := collectors.Analyze()

	if len(system.Objects) != 2 {
		t.Fatalf("objects=%d want 2", len(system.Objects))
	}

	found := false

	for _, evidence := range system.Objects[0].Evidence {
		if evidence.Type == "package_identity" &&
			evidence.Value == "rpm:tool:0:1.0-1:x86_64" {
			found = true
		}
	}

	if !found {
		t.Fatalf(
			"missing package correlation: %+v",
			system.Objects[0].Evidence,
		)
	}
}
