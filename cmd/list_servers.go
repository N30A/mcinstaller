package cmd

import (
	"fmt"

	"github.com/n30a/mcinstaller/servers"
	"github.com/spf13/cobra"
)

var listServersCmd = &cobra.Command{
	Use:                   "servers",
	Short:                 "List supported servers",
	Run:                   listServers,
	DisableFlagsInUseLine: true,
}

func listServers(cmd *cobra.Command, args []string) {
	fmt.Println("Listing all supported servers...")

	for server := servers.Vanilla; server <= servers.Forge; server++ {
		fmt.Print(server, " ")
	}
	fmt.Println()
}
