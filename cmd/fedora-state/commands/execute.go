package commands

func Execute(
	name string,
	runners map[string]Runner,
) error {
	command := CommandFromArgs([]string{name})

	return Run(command, runners)
}
