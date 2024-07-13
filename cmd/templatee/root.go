package template

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "template [template]",
	Short: "Creates/Update a template",
	Long:  `Creates/Update a template from a project`,
	Aliases: []string{
		"t",
		"temp",
		"template",
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
