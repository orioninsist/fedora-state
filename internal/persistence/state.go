package persistence

import (
	"encoding/json"
	"os"

	"fedora-state/internal/model"
)

func SaveSystem(
	path string,
	system model.System,
) error {

	content, err := json.MarshalIndent(
		system,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		content,
		0644,
	)
}

func LoadSystem(
	path string,
) (model.System, error) {

	content, err := os.ReadFile(path)

	if err != nil {
		return model.System{}, err
	}

	var system model.System

	if err := json.Unmarshal(content, &system); err != nil {
		return model.System{}, err
	}

	return system, nil
}
