package rpm

import (
	"os/exec"
)

type Provider struct {
	command func(string, ...string) *exec.Cmd
}

func New() Provider {
	return Provider{
		command: exec.Command,
	}
}

func (p Provider) Add(
	name string,
) error {

	return p.command(
		"dnf",
		"install",
		"-y",
		name,
	).Run()
}

func (p Provider) Remove(
	name string,
) error {

	return p.command(
		"dnf",
		"remove",
		"-y",
		name,
	).Run()
}
