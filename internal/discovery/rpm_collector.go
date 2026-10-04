package discovery

import (
	"os/exec"
	"strings"

	"fedora-state/internal/model"
)

type RPMCollector struct {
	command func(name string, args ...string) ([]byte, error)
}

func NewRPMCollector() RPMCollector {
	return RPMCollector{
		command: func(name string, args ...string) ([]byte, error) {
			return exec.Command(name, args...).Output()
		},
	}
}

func (collector RPMCollector) Collect() Collection {
	command := collector.command
	if command == nil {
		command = NewRPMCollector().command
	}

	output, err := command(
		"rpm",
		"-qa",
		"--qf",
		"%{NAME}\\t%{VERSION}\\t%{RELEASE}\\t%{ARCH}\\t%{INSTALLTIME}\\t%{VENDOR}\\t%{PACKAGER}\\t%{SOURCERPM}\\n",
	)
	if err != nil {
		return Collection{Err: err}
	}

	return Collection{Objects: parseRPMObjects(string(output))}
}

func parseRPMObjects(input string) []model.Object {
	var result []model.Object

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		fields := strings.Split(line, "\t")
		if len(fields) != 8 {
			continue
		}

		name := fields[0]
		version := fields[1]
		release := fields[2]
		architecture := fields[3]
		installTime := fields[4]
		vendor := fields[5]
		packager := fields[6]
		sourceRPM := fields[7]

		if name == "" || version == "" {
			continue
		}

		versionID := version
		if release != "" && release != "(none)" {
			versionID += "-" + release
		}

		evidence := []model.Evidence{
			{
				Type:  "package_database",
				Value: "rpm",
			},
		}

		for _, item := range []model.Evidence{
			{Type: "architecture", Value: architecture},
			{Type: "install_time", Value: installTime},
			{Type: "vendor", Value: vendor},
			{Type: "packager", Value: packager},
			{Type: "source_package", Value: sourceRPM},
		} {
			if item.Value != "" && item.Value != "0" && item.Value != "(none)" {
				evidence = append(evidence, item)
			}
		}

		result = append(result, model.Object{
			Name:     name,
			Type:     "package",
			Identity: "rpm:" + name + ":" + architecture,
			Version:  versionID,
			Evidence: evidence,
		})
	}

	return result
}
