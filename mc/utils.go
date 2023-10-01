package mc

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var SupportedServers = [4]string{"vanilla", "paper", "fabric", "forge"}

func CreateScript(server, serverDir, fileName string) error {

	var javaArgs string
	var scriptName string

	if server == "forge" {
		javaArgs = fmt.Sprintf("java -jar %s --installServer", fileName)

		if runtime.GOOS == "windows" {
			scriptName = "install.bat"
		} else {
			scriptName = "install.sh"
		}

	} else {
		javaArgs = fmt.Sprintf("java -Xms2G -Xmx2G -jar %s --nogui", fileName)

		if runtime.GOOS == "windows" {
			scriptName = "run.bat"
		} else {
			scriptName = "run.sh"
		}
	}

	path := filepath.Join(serverDir, scriptName)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, javaArgs)
	if err != nil {
		return err
	}

	return nil
}

func GetServerData(server string) (MCServer, error) {

	var dataFunc func() (MCServer, error)

	switch server {

	case SupportedServers[0]:
		dataFunc = GetVanillaData
	case SupportedServers[1]:
		dataFunc = GetPaperData
	case SupportedServers[2]:
		dataFunc = GetFabricData
	case SupportedServers[3]:
		dataFunc = GetForgeData
	default:
		return MCServer{}, fmt.Errorf("unsupported server: %s", server)
	}

	return dataFunc()
}

func InMCServer(version string, server MCServer) bool {
	for _, b := range server.Versions {
		if b.ID == version {
			return true
		}
	}
	return false
}
