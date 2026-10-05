package providers

import "testing"

type fakeProvider struct{}

func (fakeProvider) Add(string) error {
	return nil
}

func (fakeProvider) Remove(string) error {
	return nil
}


func TestRegistryLookup(t *testing.T) {

	registry := NewRegistry()

	registry.Register(
		"rpm",
		fakeProvider{},
	)

	_, err := registry.Lookup("rpm")

	if err != nil {
		t.Fatal(err)
	}
}
