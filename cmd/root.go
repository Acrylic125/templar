package cmd

import (
	"fmt"
	"os"

	template "acrylic125.com/templar/cmd/templatee"
	"acrylic125.com/templar/cmd/use"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "templar",
	Short: "Quickly setup a new project",
}

func init() {
	rootCmd.AddCommand(template.RootCmd)
	rootCmd.AddCommand(use.RootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
