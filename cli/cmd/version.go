package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays app version",
	Run: func(*cobra.Command, []string) {
		fmt.Println("0.0.1")
	},
}
