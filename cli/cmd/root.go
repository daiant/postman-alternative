package cmd

import (
	"cli/cmd/http"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "charlo",
		Short: "Charlo fetches data for you from the web",
	}
)

func init() {
	rootCmd.AddCommand(http.HttpCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
