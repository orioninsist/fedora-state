package commands

func DefaultRunners() map[string]Runner {
	return map[string]Runner{
		"help": func() error {
			return nil
		},
		"version": func() error {
			return nil
		},
		"plan": func() error {
			return nil
		},
		"apply": func() error {
			return nil
		},
		"manifest": func() error {
			return nil
		},
	}
}
