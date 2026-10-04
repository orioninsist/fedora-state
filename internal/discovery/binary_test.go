package discovery

import (
	"os"
	"path/filepath"
	"testing"
)

func writeExecutable(t *testing.T, path string) {
	t.Helper()

	if err := os.WriteFile(
		path,
		[]byte("#!/bin/sh\nexit 0\n"),
		0755,
	); err != nil {
		t.Fatal(err)
	}
}

func TestDiscoverBinariesUsesFirstPATHOccurrence(t *testing.T) {
	root := t.TempDir()
	binA := filepath.Join(root, "bin-a")
	binB := filepath.Join(root, "bin-b")

	if err := os.Mkdir(binA, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(binB, 0755); err != nil {
		t.Fatal(err)
	}

	targetA := filepath.Join(root, "tool-a")
	targetB := filepath.Join(root, "tool-b")

	writeExecutable(t, targetA)
	writeExecutable(t, targetB)

	if err := os.Symlink(targetA, filepath.Join(binA, "tool")); err != nil {
		t.Fatal(err)
	}

	if err := os.Symlink(targetB, filepath.Join(binB, "tool")); err != nil {
		t.Fatal(err)
	}

	t.Setenv(
		"PATH",
		binA+string(os.PathListSeparator)+binB,
	)

	got := DiscoverBinaries()

	if len(got) != 1 {
		t.Fatalf("got %d binaries, want 1: %#v", len(got), got)
	}

	if got[0].Path != filepath.Join(binA, "tool") {
		t.Errorf("path=%q, want first PATH occurrence", got[0].Path)
	}

	if got[0].RealPath != targetA {
		t.Errorf("real_path=%q, want %q", got[0].RealPath, targetA)
	}
}

func TestDiscoverBinariesPreservesDifferentNamesForSameTarget(t *testing.T) {
	root := t.TempDir()
	bin := filepath.Join(root, "bin")

	if err := os.Mkdir(bin, 0755); err != nil {
		t.Fatal(err)
	}

	target := filepath.Join(root, "real-tool")
	writeExecutable(t, target)

	for _, name := range []string{"foo", "bar"} {
		if err := os.Symlink(target, filepath.Join(bin, name)); err != nil {
			t.Fatal(err)
		}
	}

	t.Setenv("PATH", bin)

	got := DiscoverBinaries()

	if len(got) != 2 {
		t.Fatalf("got %d binaries, want 2: %#v", len(got), got)
	}
}

func TestDiscoverBinariesIgnoresNonExecutableFiles(t *testing.T) {
	dir := t.TempDir()

	if err := os.WriteFile(
		filepath.Join(dir, "not-executable"),
		[]byte("data"),
		0644,
	); err != nil {
		t.Fatal(err)
	}

	t.Setenv("PATH", dir)

	got := DiscoverBinaries()

	if len(got) != 0 {
		t.Fatalf("got unexpected binaries: %#v", got)
	}
}
