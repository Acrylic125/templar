package use

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "use [template] [-ng]",
	Short: "Templatise using the specified template",
	Long: `Creates a project from a template
	
Flags:
  -ng, --no-git    Do not initialise a git repository`,
	Aliases: []string{
		"u",
		"use",
		"apply",
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
