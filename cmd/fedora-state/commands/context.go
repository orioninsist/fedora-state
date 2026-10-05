package commands

import "fedora-state/internal/model"

type Context struct {
	Args   []string
	Home   string
	System model.System
}
