package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/njordice/mcinstaller/mc"
	"github.com/njordice/mcinstaller/utils"
)

const Usage = `
mcinstaller - minecraft server installer

Description:
  mcinstaller is a command-line tool for installing minecraft servers. It allows you to list supported servers, list available server versions, and install minecraft servers with ease.

Usage:
  mcinstaller <command> [options]

Commands:
  list servers
    List all supported minecraft servers.

  list versions <server>
    List all supported versions for the specified server.

  install <server> <version> <server-dir>
    Install a minecraft server with the supplied information.

  help
    Show this help message.

Examples:
  List all supported servers:
    $ mcinstaller list servers

  List available versions for a server:
    $ mcinstaller list versions vanilla

  Install a minecraft server:
    $ mcinstaller install vanilla 1.17.1 /path/to/server-directory

Note:
  - The <server> argument should be a valid server. (e.g. vanilla or paper).
  - The <version> argument should be a valid version number for the selected server.
  - The <server-dir> argument should be the path where you want to install the server.

`

func ListSubcommand(args []string) {

	if len(args) <= 0 {
		log.Fatalln("Too few options for 'list' - did you mean 'list [servers|versions <server>]'")
	}

	subcommand := args[0]

	if subcommand == "servers" {
		if len(args) != 1 {
			log.Fatalln("Invalid options for 'list' - did you mean 'list servers'")
		}

		log.Println("Listing all supported minecraft servers")
		for _, v := range mc.SupportedServers {
			fmt.Print(v + " ")
		}

	} else if subcommand == "versions" {
		if len(args) != 2 {
			log.Fatalln("Invalid options for 'list' - did you mean 'list versions <server>'")
		}

		server := args[1]

		log.Printf("Listing all supported versions for '%s' - please wait\n", args[1])

		serverData, err := mc.GetServerData(server)
		if err != nil {
			log.Fatalln(err)
		}

		for _, version := range serverData.Versions {
			fmt.Print(version.ID + " ")
		}

	} else {
		log.Fatalln("Invalid option for 'list' - did you mean 'list [servers|versions <server>]'")

	}
}

func InstallSubcommand(args []string) {
	if len(args) != 3 {
		log.Fatalln("Invalid options for 'install' - syntax 'install <server> <version> <server-dir>'")
	}

	server := args[0]
	version := args[1]
	serverDir := args[2]

	log.Println("Fetching server data...")

	serverData, err := mc.GetServerData(server)
	if err != nil {
		log.Fatalln(err)
	}

	if !mc.InMCServer(version, serverData) {
		log.Fatalf("unsupported version: %s\n", version)
	}

	var url string

	for _, v := range serverData.Versions {
		if version == v.ID {
			url = v.URL
			break
		}
	}

	log.Println("Successfully fetched server data")

	_, err = os.Stat(serverDir)

	if os.IsNotExist(err) {
		log.Printf("Directory '%s' does not exist, creating it\n", serverDir)
		os.MkdirAll(serverDir, os.ModePerm)

	} else if err != nil {
		log.Fatalf("Error checking path: %v\n", err)

	}

	var fileName string

	if server == mc.SupportedServers[3] {
		fileName = fmt.Sprintf("%s-installer-%s.jar", server, version)
	} else {
		fileName = fmt.Sprintf("%s-%s.jar", server, version)
	}

	filePath, _ := filepath.Abs(filepath.Join(serverDir, fileName))

	log.Println("Starting download...")
	err = utils.DownloadFile(url, filePath)
	if err != nil {
		log.Fatalln(err)

	}
	log.Printf("Successfully downloaded '%s' version '%s' at '%s'\n", server, version, filePath)

	log.Println("Creating server script")
	err = mc.CreateScript(server, serverDir, fileName)
	if err != nil {
		log.Fatalln(err)

	}

	log.Println("Installation finished, have fun playing!")
}
