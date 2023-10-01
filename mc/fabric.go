package mc

import (
	"fmt"

	"github.com/njordice/mcinstaller/utils"
)

func GetFabricData() (MCServer, error) {

	fabric := SetDefault("fabric")

	manifest, err := fetchFabricManifest()
	if err != nil {
		return MCServer{}, nil
	}

	for _, version := range manifest.Versions {

		var latestLoader, latestInstaller string

		for _, loaderIndex := range manifest.Loaders {
			if loaderIndex.Stable {
				latestLoader = loaderIndex.Version
			}
		}

		for _, installerIndex := range manifest.Installers {
			if installerIndex.Stable {
				latestInstaller = installerIndex.Version
			}
		}

		url := fmt.Sprintf("https://meta.fabricmc.net/v2/versions/loader/%s/%s/%s/server/jar", version.ID, latestLoader, latestInstaller)

		newVersion := MCVersion{
			ID:  version.ID,
			URL: url,
		}

		fabric.Versions = append(fabric.Versions, newVersion)
	}

	return fabric, nil
}

func fetchFabricManifest() (FabricManifest, error) {
	response, err := utils.GetRequest("https://meta.fabricmc.net/v2/versions/")
	if err != nil {
		return FabricManifest{}, err
	}
	defer response.Body.Close()

	var manifest FabricManifest
	err = utils.DecodeJSON(response.Body, &manifest)
	if err != nil {
		return FabricManifest{}, err
	}

	return manifest, nil
}
