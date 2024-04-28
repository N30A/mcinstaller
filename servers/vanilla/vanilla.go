package servers

import (
	"fmt"

	"github.com/n30a/mcinstaller/helpers"
)

const (
	versionManifestURL     = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
	firstVersionWithServer = "1.2.5"
)

type Vanilla struct{}

func (v *Vanilla) Versions() ([]string, error) {
	versionManifest, err := fetchVersionManifest()
	if err != nil {
		return nil, err
	}

	return versionManifest.toStringSlice(), nil
}

func (v *Vanilla) DownloadURL(version string) (string, error) {
	versionManifest, err := fetchVersionManifest()
	if err != nil {
		return "", err
	}

	index := versionManifest.indexOf(version)
	if index == -1 {
		return "", fmt.Errorf("the specified version was not found: %s", version)
	}

	versionURL := versionManifest.Versions[index].URL

	versionPackage, err := fetchVersionPackage(versionURL)
	if err != nil {
		return "", err
	}

	if versionPackage.Downloads.Server.URL == "" {
		return "", fmt.Errorf("the specified version does not have a server: %s", version)
	}

	return versionPackage.Downloads.Server.URL, nil
}

func fetchVersionManifest() (*versionManifest, error) {
	var versionManifest versionManifest

	err := helpers.FetchAndDecodeJSON(versionManifestURL, &versionManifest)
	if err != nil {
		return nil, err
	}

	index := versionManifest.indexOf(firstVersionWithServer) + 1
	versionManifest.Versions = versionManifest.Versions[:index]

	return &versionManifest, nil
}

func fetchVersionPackage(url string) (*versionPackage, error) {
	var versionPackage versionPackage

	err := helpers.FetchAndDecodeJSON(url, &versionPackage)
	if err != nil {
		return nil, err
	}

	return &versionPackage, nil
}
