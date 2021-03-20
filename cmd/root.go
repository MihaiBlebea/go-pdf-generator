package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var rootCmd = &cobra.Command{
	Use:          "go-diploma",
	Version:      "1.0.0",
	Short:        "CLI application to interact with the platform blog",
	Long:         "CLI application to interact with the platform blog",
	SilenceUsage: true,
}

// Execute _
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
