package discovery

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"fedora-state/internal/model"
)

func System() model.System {
	return model.System{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		Metadata:     metadata(),
	}
}

func metadata() model.SystemMetadata {
	hostname, _ := os.Hostname()

	kernel := ""

	if out, err := exec.Command("uname", "-r").Output(); err == nil {
		kernel = strings.TrimSpace(string(out))
	}

	release := osRelease("/etc/os-release")

	return model.SystemMetadata{
		Distribution:   release["NAME"],
		DistributionID: release["ID"],
		Version:        release["VERSION"],
		VersionID:      release["VERSION_ID"],
		Kernel:         kernel,
		Hostname:       hostname,
	}
}

func osRelease(path string) map[string]string {
	result := map[string]string{}

	file, err := os.Open(path)
	if err != nil {
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		result[strings.TrimSpace(key)] = strings.Trim(
			strings.TrimSpace(value),
			`"'`,
		)
	}

	return result
}
