package discovery

import (
	"os/exec"
	"strings"

	"fedora-state/internal/model"
)

type Resolver interface {
	Resolve([]model.Object) []model.Object
}

type RPMOwnershipResolver struct {
	command func(name string, args ...string) ([]byte, error)
}

func NewRPMOwnershipResolver() RPMOwnershipResolver {
	return RPMOwnershipResolver{
		command: func(name string, args ...string) ([]byte, error) {
			return exec.Command(name, args...).Output()
		},
	}
}

func (resolver RPMOwnershipResolver) Resolve(objects []model.Object) []model.Object {
	wanted := make(map[string]struct{})
	for _, object := range objects {
		if object.Type == "executable" && object.Location != "" {
			wanted[object.Location] = struct{}{}
		}
	}
	if len(wanted) == 0 {
		return append([]model.Object(nil), objects...)
	}

	command := resolver.command
	if command == nil {
		command = NewRPMOwnershipResolver().command
	}

	output, err := command(
		"rpm",
		"-qa",
		"--qf",
		"@@PKG@@\t%{NAME}-%{VERSION}-%{RELEASE}.%{ARCH}\n[%{FILENAMES}\n]",
	)
	if err != nil {
		return append([]model.Object(nil), objects...)
	}

	owners := make(map[string]string)
	owner := ""
	for _, line := range strings.Split(string(output), "\n") {
		if strings.HasPrefix(line, "@@PKG@@\t") {
			owner = strings.TrimPrefix(line, "@@PKG@@\t")
			continue
		}
		if owner == "" {
			continue
		}
		if _, ok := wanted[line]; ok {
			owners[line] = owner
		}
	}

	result := make([]model.Object, 0, len(objects))
	for _, object := range objects {
		if owner, ok := owners[object.Location]; ok {
			object.Evidence = append(object.Evidence, model.Evidence{
				Type:  "package_owner",
				Value: owner,
			})
		}
		result = append(result, object)
	}

	return result
}
