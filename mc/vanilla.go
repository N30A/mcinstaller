package mc

import (
	"sync"
	"time"

	"github.com/njordice/mcinstaller/utils"
)

func GetVanillaData() (MCServer, error) {

	vanilla := SetDefault("vanilla")

	manifest, err := fetchVanillaManifest()
	if err != nil {
		return MCServer{}, nil
	}

	dataChannel := make(chan MCVersion)
	var wg sync.WaitGroup

	for _, version := range manifest.Versions {
		wg.Add(1)
		time.Sleep(time.Microsecond)

		go func(url string) {
			defer wg.Done()

			data, err := fetchVanillaData(url)
			if err != nil {
				return
			}

			if data.Downloads.Server.URL != "" {
				newVersion := MCVersion{
					ID:  data.ID,
					URL: data.Downloads.Server.URL,
				}

				dataChannel <- newVersion
			}

		}(version.URL)
	}

	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	for newVersion := range dataChannel {
		vanilla.Versions = append(vanilla.Versions, newVersion)
	}

	return vanilla, nil
}

func fetchVanillaManifest() (VanillaManifest, error) {
	response, err := utils.GetRequest("https://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		return VanillaManifest{}, err
	}
	defer response.Body.Close()

	var manifest VanillaManifest
	err = utils.DecodeJSON(response.Body, &manifest)
	if err != nil {
		return VanillaManifest{}, err
	}

	return manifest, nil
}

func fetchVanillaData(url string) (VanillaData, error) {
	versionResponse, err := utils.GetRequest(url)
	if err != nil {
		return VanillaData{}, err
	}
	defer versionResponse.Body.Close()

	var data VanillaData
	err = utils.DecodeJSON(versionResponse.Body, &data)
	if err != nil {
		return VanillaData{}, err
	}

	return data, nil
}
