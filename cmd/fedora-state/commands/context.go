package commands

import (
	"fedora-state/internal/config"
	"fedora-state/internal/model"
)

type Context struct {
	Config config.Config
	Args   []string
	Home   string
	System model.System
}
