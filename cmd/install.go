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
	server, version, path := args[0], args[1], args[2]

	serverType, err := servers.NewServer(server)
	if err != nil {
		return err
	}

	fmt.Printf("Fetching %s server download url for version %s\n", server, version)
	url, err := serverType.DownloadURL(version)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist, creating it\n", path)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	var fileName string
	var suffix string

	if server == servers.SupportedServers[3] {
		suffix = "-installer"
	}

	fileName = fmt.Sprintf("%s-%s%s.jar", server, version, suffix)
	filePath := filepath.Join(path, fileName)

	fmt.Printf("Downloading %s version %s to %s\n", server, version, filePath)
	err = helpers.DownloadFile(url, filePath)
	if err != nil {
		return err
	}

	return nil
}
