package rpm

import (
	"errors"
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestRPMOwnershipResolverAddsPackageEvidence(t *testing.T) {
	resolver := RPMOwnershipResolver{
		command: func(name string, args ...string) ([]byte, error) {
			return []byte("@@PKG@@\tbash\t0\t5.3.9\t3.fc44\tx86_64\n/usr/bin/bash\n"), nil
		},
	}

	input := []model.Object{
		{
			Name:     "bash",
			Type:     "executable",
			Location: "/usr/bin/bash",
		},
	}

	got := resolver.Resolve(input)

	want := []model.Object{
		{
			Name:     "bash",
			Type:     "executable",
			Location: "/usr/bin/bash",
			Evidence: []model.Evidence{
				{
					Type:  "package_identity",
					Value: "rpm:bash:0:5.3.9-3.fc44:x86_64",
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("objects differ\n got: %#v\nwant: %#v", got, want)
	}
}

func TestRPMOwnershipResolverKeepsObjectWhenRPMFails(t *testing.T) {
	resolver := RPMOwnershipResolver{
		command: func(string, ...string) ([]byte, error) {
			return nil, errors.New("not owned")
		},
	}

	input := []model.Object{
		{
			Name:     "tool",
			Type:     "executable",
			Location: "/unknown/tool",
		},
	}

	got := resolver.Resolve(input)

	if !reflect.DeepEqual(got, input) {
		t.Fatalf("object changed\n got: %#v\nwant: %#v", got, input)
	}
}

func TestRPMOwnershipResolverIgnoresNonExecutableObjects(t *testing.T) {
	called := false

	resolver := RPMOwnershipResolver{
		command: func(string, ...string) ([]byte, error) {
			called = true
			return nil, nil
		},
	}

	input := []model.Object{
		{
			Name: "bash",
			Type: "package",
		},
	}

	got := resolver.Resolve(input)

	if called {
		t.Fatal("rpm command called for non executable object")
	}

	if !reflect.DeepEqual(got, input) {
		t.Fatalf("objects changed\n got: %#v\nwant: %#v", got, input)
	}
}

func TestRPMOwnershipResolverQueriesDatabaseOnce(t *testing.T) {
	calls := 0
	resolver := RPMOwnershipResolver{
		command: func(string, ...string) ([]byte, error) {
			calls++
			return []byte("@@PKG@@\tbash\t0\t5.3.9\t3.fc44\tx86_64\n/usr/bin/bash\n/etc/bashrc\n@@PKG@@\tcoreutils\t0\t9.10\t5.fc44\tx86_64\n/usr/bin/ls\n/usr/bin/cp\n"), nil
		},
	}

	input := []model.Object{
		{Name: "bash", Type: "executable", Location: "/usr/bin/bash"},
		{Name: "ls", Type: "executable", Location: "/usr/bin/ls"},
	}

	got := resolver.Resolve(input)

	if calls != 1 {
		t.Fatalf("rpm command called %d times, want 1", calls)
	}
	if len(got[0].Evidence) != 1 || got[0].Evidence[0].Type != "package_identity" || got[0].Evidence[0].Value != "rpm:bash:0:5.3.9-3.fc44:x86_64" {
		t.Fatalf("bash ownership = %#v", got[0].Evidence)
	}
	if len(got[1].Evidence) != 1 || got[1].Evidence[0].Type != "package_identity" || got[1].Evidence[0].Value != "rpm:coreutils:0:9.10-5.fc44:x86_64" {
		t.Fatalf("ls ownership = %#v", got[1].Evidence)
	}
}
