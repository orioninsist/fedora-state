package manifest

import (
	"sort"

	"fedora-state/internal/model"
)

func FromSystem(system model.System) Manifest {
	var entries []Entry

	for _, object := range system.Objects {
		if object.Type != "package" {
			continue
		}

		if object.Identity == "" {
			continue
		}

		entries = append(entries, Entry{
			Identity: object.Identity,
			Name:     object.Name,
			Type:     object.Type,
			Version:  object.Version,
			Source:   object.Source,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Identity < entries[j].Identity
	})

	return Manifest{
		Entries: entries,
	}
}
