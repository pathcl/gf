package app

import (
	"github.com/pathcl/gf/command"
)

func init() {
	application.AddCommand(command.Version)
	application.AddCommand(command.Search)
}
