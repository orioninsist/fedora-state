package rpm

import (
	"os/exec"
	"strings"

	"fedora-state/internal/discovery"
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

func (collector RPMCollector) Collect() discovery.Collection {
	command := collector.command
	if command == nil {
		command = NewRPMCollector().command
	}

	output, err := command(
		"rpm",
		"-qa",
		"--qf",
		"%{NAME}\\t%{EPOCHNUM}\\t%{VERSION}\\t%{RELEASE}\\t%{ARCH}\\t%{INSTALLTIME}\\t%{VENDOR}\\t%{PACKAGER}\\t%{SOURCERPM}\\n",
	)
	if err != nil {
		return discovery.Collection{Err: err}
	}

	return discovery.Collection{Objects: parseRPMObjects(string(output))}
}

func parseRPMObjects(input string) []model.Object {
	var result []model.Object

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		fields := strings.Split(line, "\t")
		if len(fields) != 9 {
			continue
		}

		name := fields[0]
		epoch := fields[1]
		version := fields[2]
		release := fields[3]
		architecture := fields[4]
		installTime := fields[5]
		vendor := fields[6]
		packager := fields[7]
		sourceRPM := fields[8]

		if name == "" || version == "" {
			continue
		}

		versionID := version
		if release != "" && release != "(none)" {
			versionID += "-" + release
		}
		if epoch != "" && epoch != "0" && epoch != "(none)" {
			versionID = epoch + ":" + versionID
		}

		identityEpoch := epoch
		if identityEpoch == "" || identityEpoch == "(none)" {
			identityEpoch = "0"
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
			Identity: "rpm:" + name + ":" + identityEpoch + ":" + version + "-" + release + ":" + architecture,
			Version:  versionID,
			Evidence: evidence,
		})
	}

	return result
}
