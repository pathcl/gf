package app

import (
	"fmt"
	"os"

	"github.com/pathcl/gf/version"

	"github.com/spf13/cobra"
)

var (
	defaults    map[string]interface{}
	application = &cobra.Command{
		Use:     version.ExecutableName(),
		Short:   version.ShortDescription(),
		Long:    version.LongDescription(),
		Version: version.VersionTemplate(),
	}
)

func Run() {
	if err := application.Execute(); err != nil {
		fmt.Errorf("Error:", err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Built with <3")
}
