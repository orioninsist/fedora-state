package commands

type Runner func() error

func Run(name string, runners map[string]Runner) error {
	command := Resolve(name)

	if runner, ok := runners[command]; ok {
		return runner()
	}

	return nil
}
