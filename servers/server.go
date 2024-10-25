package servers

import (
	"fmt"
	"strings"

	fabric "github.com/n30a/mcinstaller/servers/fabric"
	forge "github.com/n30a/mcinstaller/servers/forge"
	paper "github.com/n30a/mcinstaller/servers/paper"
	vanilla "github.com/n30a/mcinstaller/servers/vanilla"
)

type SupportedServer int

const (
	Vanilla SupportedServer = iota
	Paper
	Fabric
	Forge
)

func (s SupportedServer) String() string {
	switch s {
	case Vanilla:
		return "Vanilla"
	case Paper:
		return "Paper"
	case Fabric:
		return "Fabric"
	case Forge:
		return "Forge"
	default:
		return "Unknown"
	}
}

func ParseServerType(server string) (SupportedServer, error) {
	server = strings.ToLower(server)

	switch server {
	case "vanilla":
		return Vanilla, nil
	case "paper":
		return Paper, nil
	case "fabric":
		return Fabric, nil
	case "forge":
		return Forge, nil
	default:
		return -1, fmt.Errorf("unknown server type: %s", server)
	}
}

type Server interface {
	Versions() ([]string, error)
	DownloadURL(version string) (string, error)
}

func ServerFactory(server SupportedServer) (Server, error) {
	switch server {

	case Vanilla:
		return &vanilla.Vanilla{}, nil

	case Paper:
		return &paper.Paper{}, nil

	case Fabric:
		return &fabric.Fabric{}, nil

	case Forge:
		return &forge.Forge{}, nil

	default:
		return nil, fmt.Errorf("unsupported server '%s'", server)
	}
}
