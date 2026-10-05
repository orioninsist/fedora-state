package persistence

import (
	"encoding/json"
	"os"

	"fedora-state/internal/plan"
)

func SavePlan(
	path string,
	value plan.Plan,
) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		data,
		0644,
	)
}

func LoadPlan(
	path string,
) (plan.Plan, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return plan.Plan{}, err
	}

	var value plan.Plan

	if err := json.Unmarshal(data, &value); err != nil {
		return plan.Plan{}, err
	}

	return value, nil
}
