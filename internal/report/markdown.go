package report

import (
	"fmt"
	"io"

	"fedora-state/internal/model"
)

func WriteMarkdown(
	writer io.Writer,
	system model.System,
) error {
	if _, err := fmt.Fprintln(
		writer,
		"# Fedora State Report",
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		writer,
		"## System\n\n- OS: %s\n- Architecture: %s\n- Distribution: %s\n- Objects: %d\n- Diagnostics: %d\n\n",
		system.OS,
		system.Architecture,
		system.Metadata.Distribution,
		len(system.Objects),
		len(system.Diagnostics),
	); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(
		writer,
		"## Objects",
	); err != nil {
		return err
	}

	for _, object := range system.Objects {
		if _, err := fmt.Fprintf(
			writer,
			"- `%s` (%s)\n",
			object.Name,
			object.Type,
		); err != nil {
			return err
		}
	}

	return nil
}
