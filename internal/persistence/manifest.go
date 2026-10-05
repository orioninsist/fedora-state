package persistence

import (
	"encoding/json"
	"os"

	"fedora-state/internal/manifest"
)

func LoadManifest(path string) (manifest.Manifest, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return manifest.Manifest{}, err
	}

	var value manifest.Manifest

	if err := json.Unmarshal(content, &value); err != nil {
		return manifest.Manifest{}, err
	}

	return value, nil
}

func SaveManifest(
	path string,
	value manifest.Manifest,
) error {
	content, err := json.MarshalIndent(
		value,
		"",
		"  ",
	)
	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		append(content, '\n'),
		0644,
	)
}
