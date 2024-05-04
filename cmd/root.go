package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcinstaller",
	Short: "mcinstaller is a command-line tool for installing minecraft servers.",
	Long: `mcinstaller is a command-line tool for installing minecraft servers.
It allows you to list supported servers, list available server versions, and install minecraft servers with ease.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		listServersCmd,
		listServerVersionsCmd,
		installCmd,
	)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
