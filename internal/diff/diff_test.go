package diff

import (
	"testing"

	"fedora-state/internal/model"
)

func TestCompareDetectsAddedAndRemoved(t *testing.T) {

	old := model.System{
		Objects: []model.Object{
			{
				Identity: "rpm:bash",
				Name: "bash",
			},
		},
	}

	current := model.System{
		Objects: []model.Object{
			{
				Identity: "rpm:neovim",
				Name: "neovim",
			},
		},
	}

	result := Compare(old, current)

	if len(result.Added) != 1 {
		t.Fatalf(
			"expected one added object",
		)
	}

	if len(result.Removed) != 1 {
		t.Fatalf(
			"expected one removed object",
		)
	}
}
