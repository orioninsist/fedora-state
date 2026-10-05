package commands

func CommandFromArgs(args []string) string {
	if len(args) == 0 {
		return Help
	}

	switch args[0] {
	case "text":
		return "text"
	case "json":
		return "json"
	default:
		return Resolve(args[0])
	}
}
