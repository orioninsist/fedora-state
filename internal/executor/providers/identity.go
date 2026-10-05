package providers

import "strings"

type Identity struct {
	Provider string
	Value    string
}

func Parse(value string) Identity {
	parts := strings.SplitN(
		value,
		":",
		2,
	)

	if len(parts) != 2 {
		return Identity{
			Value: value,
		}
	}

	return Identity{
		Provider: parts[0],
		Value:    parts[1],
	}
}
