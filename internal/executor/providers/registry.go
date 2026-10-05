package providers

import "fmt"

type Handler interface {
	Add(string) error
	Remove(string) error
}

type Registry struct {
	values map[string]Handler
}

func NewRegistry() Registry {
	return Registry{
		values: map[string]Handler{},
	}
}

func (r Registry) Register(
	name string,
	handler Handler,
) {
	r.values[name] = handler
}

func (r Registry) Lookup(
	name string,
) (Handler, error) {

	value, ok := r.values[name]

	if !ok {
		return nil, fmt.Errorf(
			"unknown provider: %s",
			name,
		)
	}

	return value, nil
}
