package state

import (
	"encoding/json"
	"os"
	"time"

	"fedora-state/internal/model"
)

type SystemState struct {
	GeneratedAt time.Time `json:"generated_at"`

	OS string `json:"os"`

	Sources []model.Source `json:"sources"`

	Packages []model.Package `json:"packages"`
}

func Save(
	path string,
	data SystemState,
) error {

	data.GeneratedAt = time.Now()

	content, err := json.MarshalIndent(
		data,
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
