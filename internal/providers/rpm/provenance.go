package rpm

import (
	"os/exec"
	"strings"

	"fedora-state/internal/model"
)

type DNFProvenanceResolver struct {
	command func(name string, args ...string) ([]byte, error)
}

func NewDNFProvenanceResolver() DNFProvenanceResolver {
	return DNFProvenanceResolver{
		command: func(name string, args ...string) ([]byte, error) {
			return exec.Command(name, args...).Output()
		},
	}
}

func (resolver DNFProvenanceResolver) Resolve(objects []model.Object) []model.Object {
	command := resolver.command
	if command == nil {
		command = NewDNFProvenanceResolver().command
	}

	output, err := command(
		"dnf",
		"repoquery",
		"--installed",
		"--queryformat",
		"%{name}\\t%{epoch}\\t%{version}\\t%{release}\\t%{arch}\\t%{from_repo}\\n",
	)
	if err != nil {
		return append([]model.Object(nil), objects...)
	}

	repositories := make(map[string]string)
	for _, line := range strings.Split(string(output), "\n") {
		if line == "" {
			continue
		}

		fields := strings.Split(line, "\t")
		if len(fields) != 6 {
			continue
		}

		name := fields[0]
		epoch := fields[1]
		version := fields[2]
		release := fields[3]
		architecture := fields[4]
		repository := fields[5]

		if name == "" || version == "" || repository == "" {
			continue
		}
		if epoch == "" || epoch == "(none)" {
			epoch = "0"
		}

		identity := "rpm:" + name + ":" + epoch + ":" +
			version + "-" + release + ":" + architecture
		repositories[identity] = repository
	}

	result := make([]model.Object, 0, len(objects))
	for _, object := range objects {
		if object.Type == "package" {
			if repository, ok := repositories[object.Identity]; ok {
				object.Evidence = append(object.Evidence, model.Evidence{
					Type:  "repository",
					Value: repository,
				})
			}
		}
		result = append(result, object)
	}

	return result
}
