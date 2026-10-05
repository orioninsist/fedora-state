package cargo

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"fedora-state/internal/discovery"

	"fedora-state/internal/model"
)

type Collector struct {
	path string
	read func(string) ([]byte, error)
}

func NewCollector(path string) Collector {
	return Collector{path: path, read: os.ReadFile}
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

	objects, err := parseObjects(input)
	if err != nil {
		return discovery.Collection{Err: err}
	}
	return discovery.Collection{Objects: objects}
}

type metadata struct {
	Installs map[string]install `json:"installs"`
}

type install struct {
	Bins   []string `json:"bins"`
	Target string   `json:"target"`
}

func parseObjects(input []byte) ([]model.Object, error) {
	var data metadata
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	var objects []model.Object
	for key, item := range data.Installs {
		name, version, source, ok := parseKey(key)
		if !ok {
			continue
		}

		evidence := []model.Evidence{{Type: "package_database", Value: "cargo"}}
		if source != "" {
			evidence = append(evidence, model.Evidence{Type: "source", Value: source})
		}
		if item.Target != "" {
			evidence = append(evidence, model.Evidence{Type: "target", Value: item.Target})
		}
		for _, bin := range item.Bins {
			evidence = append(evidence, model.Evidence{Type: "executable_name", Value: bin})
		}

		objects = append(objects, model.Object{
			Name:     name,
			Type:     "package",
			Identity: "cargo:" + name + ":" + version + ":" + source,
			Version:  version,
			Evidence: evidence,
		})
	}

	return objects, nil
}

func parseKey(key string) (string, string, string, bool) {
	open := strings.Index(key, " (")
	if open < 0 || !strings.HasSuffix(key, ")") {
		return "", "", "", false
	}

	fields := strings.Fields(key[:open])
	if len(fields) != 2 {
		return "", "", "", false
	}

	return fields[0], fields[1], key[open+2 : len(key)-1], true
}
