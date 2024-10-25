package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/n30a/mcinstaller/helpers"
	"github.com/n30a/mcinstaller/servers"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:                   "install <server> <version> <path>",
	Short:                 "Install a minecraft server",
	Args:                  cobra.ExactArgs(3),
	RunE:                  install,
	DisableFlagsInUseLine: true,
}

func install(cmd *cobra.Command, args []string) error {
	serverArg, versionArg, pathArg := args[0], args[1], args[2]

	serverType, err := servers.ParseServerType(serverArg)
	if err != nil {
		return err
	}

	server, err := servers.ServerFactory(serverType)
	if err != nil {
		return err
	}

	fmt.Printf("Fetching %s server download url for version %s\n", serverType, versionArg)
	url, err := server.DownloadURL(versionArg)
	if err != nil {
		return err
	}

	if _, err := os.Stat(pathArg); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist, creating it\n", pathArg)
		err = os.MkdirAll(pathArg, os.ModePerm)
		if err != nil {
			return err
		}
	}

	var fileName string
	var suffix string

	if serverType == servers.Forge {
		suffix = "-installer"
	}

	fileName = fmt.Sprintf("%s-%s%s.jar", serverType, versionArg, suffix)
	filePath := filepath.Join(pathArg, fileName)

	fmt.Printf("Downloading %s version %s to %s\n", serverType, versionArg, filePath)
	err = helpers.DownloadFile(url, filePath)
	if err != nil {
		return err
	}

	return nil
}
