package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"fedora-state/internal/state"
)

func State(data state.SystemState) (string, error) {
	content, err := json.Marshal(struct {
		OS       string `json:"os"`
		Sources  any    `json:"sources"`
		Packages any    `json:"packages"`
	}{
		OS:       data.OS,
		Sources:  data.Sources,
		Packages: data.Packages,
	})
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(content)

	return hex.EncodeToString(sum[:]), nil
}
