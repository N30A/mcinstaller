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
	serverType, err := servers.ParseServerType(args[0])
	if err != nil {
		return err
	}

	server, err := servers.ServerFactory(serverType)
	if err != nil {
		return err
	}

	versions, err := server.Versions()
	if err != nil {
		return err
	}

	fmt.Printf("Listing all versions for %s...\n", serverType)
	for _, version := range versions {
		fmt.Println(version)
	}

	return nil
}
