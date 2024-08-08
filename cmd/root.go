package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IPTracker CLI app",
		Long:  `IPTracker CLI app.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
