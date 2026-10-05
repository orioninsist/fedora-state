package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"

	"fedora-state/internal/model"
)

func State(system model.System) (string, error) {
	normalized := system

	normalized.Objects = append([]model.Object(nil), system.Objects...)
	for index := range normalized.Objects {
		normalized.Objects[index].Evidence = append(
			[]model.Evidence(nil),
			normalized.Objects[index].Evidence...,
		)

		sort.Slice(normalized.Objects[index].Evidence, func(i, j int) bool {
			left := normalized.Objects[index].Evidence[i]
			right := normalized.Objects[index].Evidence[j]

			if left.Type != right.Type {
				return left.Type < right.Type
			}
			if left.Path != right.Path {
				return left.Path < right.Path
			}
			return left.Value < right.Value
		})
	}

	sort.Slice(normalized.Objects, func(i, j int) bool {
		left := normalized.Objects[i]
		right := normalized.Objects[j]

		if left.Type != right.Type {
			return left.Type < right.Type
		}
		if left.Identity != right.Identity {
			return left.Identity < right.Identity
		}
		if left.Name != right.Name {
			return left.Name < right.Name
		}
		return left.Location < right.Location
	})

	normalized.Diagnostics = append(
		[]model.Diagnostic(nil),
		system.Diagnostics...,
	)
	sort.Slice(normalized.Diagnostics, func(i, j int) bool {
		return normalized.Diagnostics[i].Message <
			normalized.Diagnostics[j].Message
	})

	content, err := json.Marshal(normalized)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:]), nil
}
