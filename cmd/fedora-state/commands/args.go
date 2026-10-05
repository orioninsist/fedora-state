package commands

func CommandFromArgs(args []string) string {
	if len(args) == 0 {
		return "help"
	}

	return Resolve(args[0])
}
