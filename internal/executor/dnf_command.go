package executor

import (
	"os/exec"

	"fedora-state/internal/plan"
)

type CommandDNFBackend struct {
	command func(string, ...string) *exec.Cmd
}

func NewCommandDNFBackend() CommandDNFBackend {
	return CommandDNFBackend{
		command: exec.Command,
	}
}

func (b CommandDNFBackend) Execute(
	action plan.Action,
) error {
	switch action.Type {
	case "add":
		return nil
	case "remove":
		return nil
	default:
		return nil
	}
}
