package cmd

import (
	"fmt"

	"github.com/n30a/mcinstaller/servers"

	"github.com/spf13/cobra"
)

var listServerVersionsCmd = &cobra.Command{
	Use:                   "versions <server>",
	Short:                 "List versions for the specified server",
	Args:                  cobra.ExactArgs(1),
	RunE:                  listServerVersions,
	DisableFlagsInUseLine: true,
}

func listServerVersions(cmd *cobra.Command, args []string) error {
	server := args[0]

	serverType, err := servers.NewServer(server)
	if err != nil {
		return err
	}

	versions, err := serverType.Versions()
	if err != nil {
		return err
	}

	fmt.Printf("Listing all versions for %s...\n", server)
	for _, version := range versions {
		fmt.Print(version + " ")
	}

	return nil
}
