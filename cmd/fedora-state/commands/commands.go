package commands

type Command struct {
	Name string
}

func Available() []Command {
	return []Command{
		{Name: "plan"},
		{Name: "apply"},
		{Name: "manifest"},
		{Name: "snapshot"},
		{Name: "snapshots"},
		{Name: "diff"},
		{Name: "version"},
		{Name: "help"},
	}
}
