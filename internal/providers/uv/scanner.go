package uv

import (
	"os"
	"path/filepath"

	"fedora-state/internal/discovery"
	"fedora-state/internal/model"
)

type Scanner struct {
	root string
}

func NewScanner(root string) Scanner {
	return Scanner{
		root: root,
	}
}

func (scanner Scanner) Collect() discovery.Collection {
	entries, err := os.ReadDir(scanner.root)
	if err != nil {
		if os.IsNotExist(err) {
			return discovery.Collection{}
		}

		return discovery.Collection{Err: err}
	}

	var objects []model.Object

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		receipt := filepath.Join(
			scanner.root,
			entry.Name(),
			"uv-receipt.toml",
		)

		collector := NewCollector(receipt)
		collection := collector.Collect()

		if collection.Err != nil {
			return discovery.Collection{Err: collection.Err}
		}

		objects = append(objects, collection.Objects...)
	}

	return discovery.Collection{
		Objects: objects,
	}
}
