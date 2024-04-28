package servers

import (
	"fmt"
	"slices"

	"github.com/n30a/mcinstaller/helpers"
)

const (
	versionManifestURL = "https://meta.fabricmc.net/v2/versions"
)

type Fabric struct{}

func (f *Fabric) Versions() ([]string, error) {
	versionManifest, err := fetchVersionManifest()
	if err != nil {
		return nil, err
	}

	return versionManifest.toStringSlice(), nil
}

func (f *Fabric) DownloadURL(version string) (string, error) {
	versionManifest, err := fetchVersionManifest()
	if err != nil {
		return "", err
	}

	if !slices.Contains(versionManifest.toStringSlice(), version) {
		return "", fmt.Errorf("the specified version was not found: %s", version)
	}

	lastedLoader := versionManifest.Loader[0].Version
	lastedInstaller := versionManifest.Installer[0].Version

	url := fmt.Sprintf(
		"https://meta.fabricmc.net/v2/versions/loader/%s/%s/%s/server/jar",
		version,
		lastedLoader,
		lastedInstaller,
	)

	return url, err
}

func fetchVersionManifest() (versionManifest, error) {
	var versionManifest versionManifest

	err := helpers.FetchAndDecodeJSON(versionManifestURL, &versionManifest)
	if err != nil {
		return versionManifest, err
	}

	return versionManifest, nil
}
