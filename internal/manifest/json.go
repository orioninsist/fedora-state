package manifest

import (
	"encoding/json"
	"io"
)

func WriteJSON(writer io.Writer, manifest Manifest) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	return encoder.Encode(manifest)
}
