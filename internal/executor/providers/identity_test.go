package providers

import "testing"

func TestParseIdentity(t *testing.T) {

	got := Parse(
		"rpm:bash:0:5.3.0-2.fc44:x86_64",
	)

	if got.Provider != "rpm" {
		t.Fatalf(
			"provider=%s",
			got.Provider,
		)
	}

	if got.Value == "" {
		t.Fatal("missing value")
	}
}
