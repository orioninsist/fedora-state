package executor

import (
	"fmt"
	"os/exec"
	"strings"

	"fedora-state/internal/plan"
)

type RealDNFBackend struct {
	command func(string, ...string) *exec.Cmd
}

func NewRealDNFBackend() RealDNFBackend {
	return RealDNFBackend{
		command: exec.Command,
	}
}

func (b RealDNFBackend) Execute(
	action plan.Action,
) error {
	switch action.Type {
	case "add":
		return b.run("dnf", "install", "-y", action.Identity)
	case "remove":
		return b.run("dnf", "remove", "-y", action.Identity)
	default:
		return fmt.Errorf(
			"unsupported action: %s",
			action.Type,
		)
	}
}

func (b RealDNFBackend) run(
	command string,
	args ...string,
) error {
	cmd := b.command(command, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"%s: %s: %w",
			command,
			strings.TrimSpace(string(output)),
			err,
		)
	}

	return nil
}
