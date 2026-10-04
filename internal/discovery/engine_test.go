package discovery

import (
	"reflect"
	"testing"
)

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
