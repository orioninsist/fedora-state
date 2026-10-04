package report

import (
	"fmt"
	"os"
	"time"

	"fedora-state/internal/state"
)

func GenerateMarkdown(
	path string,
	data state.SystemState,
) error {

	if data.GeneratedAt.IsZero() {
		data.GeneratedAt = time.Now()
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	fmt.Fprintln(
		file,
		"# Fedora System Report",
	)

	fmt.Fprintln(
		file,
	)

	fmt.Fprintf(
		file,
		"Generated: %s\n\n",
		data.GeneratedAt.Format("2006-01-02 15:04:05"),
	)

	fmt.Fprintln(
		file,
		"## System",
	)

	fmt.Fprintln(
		file,
	)

	fmt.Fprintf(
		file,
		"- OS: %s\n\n",
		data.OS,
	)

	fmt.Fprintln(
		file,
		"## Sources",
	)

	fmt.Fprintln(
		file,
	)

	fmt.Fprintln(
		file,
		"| Source | Available | Collected |",
	)

	fmt.Fprintln(
		file,
		"|---|---|---|",
	)

	for _, source := range data.Sources {

		fmt.Fprintf(
			file,
			"| %s | %t | %t |\n",
			source.Name,
			source.Available,
			source.Collected,
		)
	}

	fmt.Fprintln(
		file,
	)

	fmt.Fprintln(
		file,
		"## Packages",
	)

	fmt.Fprintln(
		file,
	)

	fmt.Fprintf(
		file,
		"Total packages: %d\n\n",
		len(data.Packages),
	)

	fmt.Fprintln(
		file,
		"| Name | Version | Release | Arch | Manager |",
	)

	fmt.Fprintln(
		file,
		"|---|---|---|---|---|",
	)

	for _, pkg := range data.Packages {

		fmt.Fprintf(
			file,
			"| %s | %s | %s | %s | %s |\n",
			pkg.Name,
			pkg.Version,
			pkg.Release,
			pkg.Arch,
			pkg.Manager,
		)
	}

	return nil
}
