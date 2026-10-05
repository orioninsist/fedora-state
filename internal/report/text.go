package report

import (
	"fmt"
	"io"

	"fedora-state/internal/model"
)

func WriteText(
	writer io.Writer,
	system model.System,
) error {
	_, err := fmt.Fprintf(
		writer,
		"os=%s architecture=%s distribution=%s objects=%d diagnostics=%d\n",
		system.OS,
		system.Architecture,
		system.Metadata.Distribution,
		len(system.Objects),
		len(system.Diagnostics),
	)

	return err
}
