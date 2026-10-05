package report

import (
	"encoding/json"
	"io"

	"fedora-state/internal/model"
)

func WriteJSON(
	writer io.Writer,
	system model.System,
) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	return encoder.Encode(system)
}
