package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"gf/version"
)

var (
	Version = &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Long:  fmt.Sprintf("version shows the version details for the %s application.", version.ApplicationName()),
		Run:   executeVersionCommand,
	}
)

func executeVersionCommand(cmd *cobra.Command, args []string) {
	fmt.Println(version.VersionTemplate())
}
