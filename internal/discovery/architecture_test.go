package discovery

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDiscoveryDoesNotDependOnLegacyPackageModel(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range files {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}

		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		source := string(content)

		for _, forbidden := range []string{
			"model.Package",
			"model.Source",
			"DiscoverPackages",
			"RPMPackages",
			"Sources()",
		} {
			if strings.Contains(source, forbidden) {
				t.Errorf(
					"%s contains legacy dependency %q",
					path,
					forbidden,
				)
			}
		}
	}
}
