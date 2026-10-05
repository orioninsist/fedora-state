package persistence

import (
	"encoding/json"
	"os"

	"fedora-state/internal/model"
)

func SaveSnapshot(
	path string,
	value model.System,
) error {

	data, err := json.MarshalIndent(
		value,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		append(data, '\n'),
		0644,
	)
}

func LoadSnapshot(
	path string,
) (model.System, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return model.System{}, err
	}

	var value model.System

	if err := json.Unmarshal(data, &value); err != nil {
		return model.System{}, err
	}

	return value, nil
}
