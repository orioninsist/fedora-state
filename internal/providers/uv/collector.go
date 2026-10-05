package uv

import (
	"errors"
	"os"
	"regexp"

	"fedora-state/internal/discovery"
	"fedora-state/internal/model"
)

type Collector struct {
	path string
	read func(string) ([]byte, error)
}

func NewCollector(path string) Collector {
	return Collector{
		path: path,
		read: os.ReadFile,
	}
}

func (collector Collector) Collect() discovery.Collection {
	read := collector.read
	if read == nil {
		read = os.ReadFile
	}

	input, err := read(collector.path)
	if errors.Is(err, os.ErrNotExist) {
		return discovery.Collection{}
	}
	if err != nil {
		return discovery.Collection{Err: err}
	}

	objects, err := parseReceipt(input)
	if err != nil {
		return discovery.Collection{Err: err}
	}

	return discovery.Collection{
		Objects: objects,
	}
}

func parseReceipt(input []byte) ([]model.Object, error) {
	text := string(input)

	requirement := regexp.MustCompile(`name = "([^"]+)".*git = "([^"]+)"`)

	match := requirement.FindStringSubmatch(text)
	if len(match) != 3 {
		return nil, errors.New("invalid uv requirement")
	}

	name := match[1]
	source := match[2]

	installPath := ""

	entrypoint := regexp.MustCompile(`install-path = "([^"]+)"`)
	if match := entrypoint.FindStringSubmatch(text); len(match) == 2 {
		installPath = match[1]
	}

	evidence := []model.Evidence{
		{
			Type:  "package_database",
			Value: "uv",
		},
		{
			Type:  "source",
			Value: source,
		},
	}

	if installPath != "" {
		evidence = append(evidence, model.Evidence{
			Type:  "binary",
			Value: installPath,
		})
	}

	return []model.Object{
		{
			Name:     name,
			Type:     "package",
			Identity: "uv:" + name + ":" + source,
			Evidence: evidence,
		},
	}, nil
}
