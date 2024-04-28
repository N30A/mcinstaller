package servers

import (
	"fmt"
	"slices"

	"github.com/n30a/mcinstaller/helpers"
)

const (
	versionManifestURL = "https://api.papermc.io/v2/projects/paper"
	versionBuildsURL   = versionManifestURL + "/versions"
)

type Paper struct{}

func (p *Paper) Versions() ([]string, error) {
	versionManifest, err := fetchVersionManifest()
	if err != nil {
		return nil, err
	}

	return versionManifest.Versions, nil
}

func (p *Paper) DownloadURL(version string) (string, error) {
	versions, err := p.Versions()
	if err != nil {
		return "", err
	}

	if !slices.Contains(versions, version) {
		return "", fmt.Errorf("the specified version was not found: %s", version)
	}

	versionBuilds, err := fetchVersionBuilds(version)
	if err != nil {
		return "", err
	}

	latestBuild := fmt.Sprint(versionBuilds.Builds[0])

	url := fmt.Sprintf("%s/%s/builds/%s/downloads/paper-%s-%s.jar",
		versionBuildsURL,
		version,
		latestBuild,
		version,
		latestBuild,
	)

	return url, nil
}

func fetchVersionManifest() (versionManifest, error) {
	var versionManifest versionManifest

	err := helpers.FetchAndDecodeJSON(versionManifestURL, &versionManifest)
	if err != nil {
		return versionManifest, err
	}

	slices.Reverse(versionManifest.Versions)

	return versionManifest, nil
}

func fetchVersionBuilds(version string) (versionBuilds, error) {
	var versionBuilds versionBuilds
	url := versionBuildsURL + "/" + version

	err := helpers.FetchAndDecodeJSON(url, &versionBuilds)
	if err != nil {
		return versionBuilds, err
	}

	slices.Reverse(versionBuilds.Builds)

	return versionBuilds, nil
}
