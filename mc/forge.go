package mc

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/njordice/mcinstaller/utils"

	"github.com/gocolly/colly/v2"
)

func GetForgeData() (MCServer, error) {

	forge := SetDefault("forge")

	manifest, err := fetchForgeManifest()
	if err != nil {
		return MCServer{}, nil
	}

	dataChannel := make(chan MCVersion)
	var wg sync.WaitGroup

	for _, version := range manifest {
		wg.Add(1)
		time.Sleep(time.Microsecond)

		go func(version string) {
			defer wg.Done()

			var url string

			c2 := colly.NewCollector()
			c2.OnHTML("a[href][title='Installer']", func(e *colly.HTMLElement) {
				url = e.Attr("href")
			})

			err := c2.Visit(fmt.Sprintf("https://files.minecraftforge.net/net/minecraftforge/forge/index_%s.html", version))
			if err != nil {
				return
			}

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
		forge.Versions = append(forge.Versions, newVersion)
	}

	return forge, nil
}

func fetchForgeManifest() ([]string, error) {
	c := colly.NewCollector()

	var output string
	c.OnHTML("a.elem-text.toggle-collapsible", func(e *colly.HTMLElement) {
		output += e.DOM.NextFiltered("ul.nav-collapsible").Text()
	})

	err := c.Visit("https://files.minecraftforge.net/net/minecraftforge/forge/")
	if err != nil {
		return nil, err
	}

	pattern := `\d+\.\d+(\.\d+)?`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(output, -1)

	blacklist := []string{"1.5.1", "1.5", "1.4.7", "1.4.6", "1.4.5", "1.4.4", "1.4.3", "1.4.2", "1.4.1", "1.4.0", "1.3.2", "1.2.5", "1.2.4", "1.2.3", "1.1"}

	var versions []string

	for _, version := range matches {
		if !utils.InSlice(version, blacklist) {
			versions = append(versions, version)
		}
	}

	return versions, nil
}
