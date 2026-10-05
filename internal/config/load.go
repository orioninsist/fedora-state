package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(
	path string,
) (Config, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var value Config

	if err := yaml.Unmarshal(data, &value); err != nil {
		return Config{}, err
	}

	return value, nil
}
