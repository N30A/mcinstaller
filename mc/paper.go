package mc

import (
	"fmt"
	"sync"
	"time"

	"github.com/njordice/mcinstaller/utils"
)

func GetPaperData() (MCServer, error) {

	paper := SetDefault("vanilla")

	manifest, err := fetchPaperManifest()
	if err != nil {
		return MCServer{}, nil
	}

	dataChannel := make(chan MCVersion)
	var wg sync.WaitGroup

	for _, version := range manifest.Versions {
		wg.Add(1)
		time.Sleep(time.Microsecond)

		go func(version string) {
			defer wg.Done()

			data, err := fetchPaperData(version)
			if err != nil {
				return
			}

			latestBuild := fmt.Sprint(data.Builds[len(data.Builds)-1])

			url := fmt.Sprintf("https://api.papermc.io/v2/projects/paper/versions/%s/builds/%s/downloads/paper-%s-%s.jar", version, latestBuild, version, latestBuild)

			newVersion := MCVersion{
				ID:  version,
				URL: url,
			}

			dataChannel <- newVersion

		}(version)
	}

	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	for newVersion := range dataChannel {
		paper.Versions = append(paper.Versions, newVersion)
	}

	return paper, nil
}

func fetchPaperManifest() (PaperManifest, error) {
	response, err := utils.GetRequest("https://api.papermc.io/v2/projects/paper")
	if err != nil {
		return PaperManifest{}, err
	}
	defer response.Body.Close()

	var manifest PaperManifest
	err = utils.DecodeJSON(response.Body, &manifest)
	if err != nil {
		return PaperManifest{}, err
	}

	for i, j := 0, len(manifest.Versions)-1; i < j; i, j = i+1, j-1 {
		manifest.Versions[i], manifest.Versions[j] = manifest.Versions[j], manifest.Versions[i]
	}

	return manifest, nil
}

func fetchPaperData(version string) (PaperData, error) {
	response, err := utils.GetRequest("https://api.papermc.io/v2/projects/paper/versions/" + version)
	if err != nil {
		return PaperData{}, err
	}
	defer response.Body.Close()

	var data PaperData
	err = utils.DecodeJSON(response.Body, &data)
	if err != nil {
		return PaperData{}, err
	}

	return data, err
}
