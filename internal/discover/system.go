package discover

import (
	"fmt"
	"os"
	"os/exec"
)

type SystemInfo struct {
	OS              string
	PackageManagers []string
}

func (s SystemInfo) String() string {

	return fmt.Sprintf(
		"OS: %s\nPackage Managers: %v",
		s.OS,
		s.PackageManagers,
	)
}

func System() (SystemInfo, error) {

	info := SystemInfo{
		OS: "unknown",
	}

	if _, err := os.Stat("/etc/fedora-release"); err == nil {
		info.OS = "fedora"
	}

	check := []string{
		"dnf",
		"flatpak",
		"cargo",
		"npm",
	}

	for _, tool := range check {

		if _, err := exec.LookPath(tool); err == nil {

			info.PackageManagers =
				append(
					info.PackageManagers,
					tool,
				)
		}
	}

	return info, nil
}
