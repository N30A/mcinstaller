package servers

import (
	"fmt"

	fabric "github.com/n30a/mcinstaller/servers/fabric"
	forge "github.com/n30a/mcinstaller/servers/forge"
	paper "github.com/n30a/mcinstaller/servers/paper"
	vanilla "github.com/n30a/mcinstaller/servers/vanilla"
)

var SupportedServers = []string{
	"vanilla",
	"paper",
	"fabric",
	"forge",
}

type Server interface {
	Versions() ([]string, error)
	DownloadURL(version string) (string, error)
}

func NewServer(server string) (Server, error) {
	var serverType Server

	switch server {

	case SupportedServers[0]:
		serverType = &vanilla.Vanilla{}

	case SupportedServers[1]:
		serverType = &paper.Paper{}

	case SupportedServers[2]:
		serverType = &fabric.Fabric{}

	case SupportedServers[3]:
		serverType = &forge.Forge{}

	default:
		return nil, fmt.Errorf("unsupported server '%s'", server)
	}

	return serverType, nil
}
