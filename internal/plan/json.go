package plan

import (
	"encoding/json"
	"io"
)

func WriteJSON(
	writer io.Writer,
	value Plan,
) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	return encoder.Encode(value)
}
