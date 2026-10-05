package commands

func Exists(name string) bool {
	for _, command := range Available() {
		if command.Name == name {
			return true
		}
	}

	return false
}
