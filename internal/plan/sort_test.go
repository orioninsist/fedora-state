package plan

import (
	"reflect"
	"testing"

	"fedora-state/internal/manifest"
)

func TestBuildPlanIsDeterministic(t *testing.T) {
	old := manifest.Manifest{}

	current := manifest.Manifest{
		Entries: []manifest.Entry{
			{
				Identity: "rpm:vim:0:1.0-1:x86_64",
				Name:     "vim",
				Type:     "package",
			},
			{
				Identity: "rpm:bash:0:1.0-1:x86_64",
				Name:     "bash",
				Type:     "package",
			},
		},
	}

	first := Build(old, current)
	second := Build(old, current)

	if !reflect.DeepEqual(first, second) {
		t.Fatalf(
			"plan differs\nfirst=%#v\nsecond=%#v",
			first,
			second,
		)
	}
}
