package commands

func Resolve(name string) string {
	if Exists(name) {
		return name
	}

	return "help"
}
