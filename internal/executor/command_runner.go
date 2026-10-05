package executor

import (
	"fmt"
	"os/exec"

	"fedora-state/internal/plan"
)

type RunnerBackend struct {
	command func(string, ...string) *exec.Cmd
}

func NewRunnerBackend() RunnerBackend {
	return RunnerBackend{
		command: exec.Command,
	}
}

func (b RunnerBackend) Execute(
	action plan.Action,
) error {
	switch action.Type {
	case "add":
		return nil
	case "remove":
		return nil
	default:
		return fmt.Errorf("unsupported action: %s", action.Type)
	}
}
