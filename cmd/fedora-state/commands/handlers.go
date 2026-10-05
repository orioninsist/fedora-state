package commands

func HandlerMap() map[string]Handler {
	return map[string]Handler{
		Help: func(Context) error {
			return nil
		},
		Version: func(Context) error {
			return nil
		},
		Plan: func(Context) error {
			return nil
		},
		Apply: func(Context) error {
			return nil
		},
		Manifest: func(Context) error {
			return nil
		},
	}
}
