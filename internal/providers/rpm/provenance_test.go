package rpm

import (
	"errors"
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestDNFProvenanceResolverAddsRepositoryEvidence(t *testing.T) {
	calls := 0
	resolver := DNFProvenanceResolver{
		command: func(name string, args ...string) ([]byte, error) {
			calls++
			return []byte(
				"bash\t0\t5.3.0\t2.fc44\tx86_64\tupdates\n" +
					"sops\t0\t3.13.3\t1\tx86_64\t@commandline\n",
			), nil
		},
	}

	input := []model.Object{
		{Name: "bash", Type: "package", Identity: "rpm:bash:0:5.3.0-2.fc44:x86_64"},
		{Name: "sops", Type: "package", Identity: "rpm:sops:0:3.13.3-1:x86_64"},
		{Name: "tool", Type: "executable", Location: "/usr/bin/tool"},
	}

	got := resolver.Resolve(input)

	want := []model.Object{
		{
			Name:     "bash",
			Type:     "package",
			Identity: "rpm:bash:0:5.3.0-2.fc44:x86_64",
			Evidence: []model.Evidence{{Type: "repository", Value: "updates"}},
		},
		{
			Name:     "sops",
			Type:     "package",
			Identity: "rpm:sops:0:3.13.3-1:x86_64",
			Evidence: []model.Evidence{{Type: "repository", Value: "@commandline"}},
		},
		{Name: "tool", Type: "executable", Location: "/usr/bin/tool"},
	}

	if calls != 1 {
		t.Fatalf("dnf command called %d times, want 1", calls)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("objects differ\n got: %#v\nwant: %#v", got, want)
	}
}

func TestDNFProvenanceResolverRequiresExactNEVRA(t *testing.T) {
	resolver := DNFProvenanceResolver{
		command: func(string, ...string) ([]byte, error) {
			return []byte("kernel\t0\t7.2.8\t200.fc44\tx86_64\tupdates\n"), nil
		},
	}

	input := []model.Object{
		{Name: "kernel", Type: "package", Identity: "rpm:kernel:0:7.2.7-200.fc44:x86_64"},
	}

	if got := resolver.Resolve(input); !reflect.DeepEqual(got, input) {
		t.Fatalf("objects changed\n got: %#v\nwant: %#v", got, input)
	}
}

func TestDNFProvenanceResolverKeepsObjectsWhenDNFFails(t *testing.T) {
	resolver := DNFProvenanceResolver{
		command: func(string, ...string) ([]byte, error) {
			return nil, errors.New("dnf unavailable")
		},
	}

	input := []model.Object{
		{Name: "bash", Type: "package", Identity: "rpm:bash:0:5.3.0-2.fc44:x86_64"},
	}

	if got := resolver.Resolve(input); !reflect.DeepEqual(got, input) {
		t.Fatalf("objects changed\n got: %#v\nwant: %#v", got, input)
	}
}
