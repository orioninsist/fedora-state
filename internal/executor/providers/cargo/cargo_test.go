package cargo

import "testing"

func TestProviderExists(t *testing.T) {

	provider := New()

	if provider.command == nil {
		t.Fatal("missing command runner")
	}
}
