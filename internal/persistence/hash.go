package persistence

import (
	"os"
	"strings"
)

func LoadHash(path string) (string, error) {

	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}

func SaveHash(
	path string,
	hash string,
) error {

	return os.WriteFile(
		path,
		[]byte(hash+"\n"),
		0644,
	)
}
