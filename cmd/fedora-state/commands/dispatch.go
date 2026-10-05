package commands

func Dispatch(
	name string,
	ctx Context,
	handlers map[string]Handler,
) error {
	command := CommandFromArgs([]string{name})

	handler, ok := handlers[command]
	if !ok {
		handler = handlers[Help]
	}

	return handler(ctx)
}
