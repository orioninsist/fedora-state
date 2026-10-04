package discovery

import (
	"bufio"
	"os/exec"
	"strings"
)

import "fedora-state/internal/model"

func RPMPackages() ([]model.Package, error) {

	cmd := exec.Command(
		"rpm",
		"-qa",
		"--qf",
		"%{NAME}|%{VERSION}|%{RELEASE}|%{ARCH}\n",
	)

	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var packages []model.Package

	scanner := bufio.NewScanner(
		strings.NewReader(string(out)),
	)

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(
			line,
			"|",
		)

		if len(parts) != 4 {
			continue
		}

		packages = append(
			packages,
			model.Package{
				Name:    parts[0],
				Version: parts[1],
				Release: parts[2],
				Arch:    parts[3],
				Manager: "rpm",
			},
		)
	}

	return packages, nil
}
