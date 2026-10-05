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
	case "markdown":
		return "markdown"
	default:
		return Resolve(args[0])
	}
}
