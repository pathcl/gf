package app

import (
	"gf/command"
)

func init() {
	application.AddCommand(command.Version)
	application.AddCommand(command.Search)
}
