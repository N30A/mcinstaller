package servers

import (
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

	versions := filterVersions(versionManifest)

	return versions, nil
}

func (v *Vanilla) DownloadURL(version string) (string, error) {
	return "", nil
}

func fetchVersionManifest() (*versionManifest, error) {
	var versionManifest versionManifest

	err := helpers.FetchAndDecodeJSON(versionManifestURL, &versionManifest)
	if err != nil {
		return nil, err
	}

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

func filterVersions(versionManifest *versionManifest) []string {
	indexOfFirstVersionWithServer := versionManifest.indexOf(firstVersionWithServer) + 1

	versions := make([]string, indexOfFirstVersionWithServer)

	for i := 0; i < len(versions); i++ {
		versions[i] = versionManifest.Versions[i].ID
	}

	return versions
}
