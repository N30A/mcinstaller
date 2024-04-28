package servers

import (
	"fmt"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/n30a/mcinstaller/helpers"
)

const (
	filesURL = "https://files.minecraftforge.net/net/minecraftforge/forge/"
)

var versionsWithNoJarFile = []string{
	"1.5.1",
	"1.5",
	"1.4.7",
	"1.4.6",
	"1.4.5",
	"1.4.4",
	"1.4.3",
	"1.4.2",
	"1.4.1",
	"1.4.0",
	"1.3.2",
	"1.2.5",
	"1.2.4",
	"1.2.3",
	"1.1",
}

type Forge struct{}

func (f *Forge) Versions() ([]string, error) {
	response, err := helpers.GetRequest(filesURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	var versions []string

	ul := doc.Find("ul.nav-collapsible")
	ul.Find("li").Each(func(i int, s *goquery.Selection) {
		version := strings.TrimSpace(s.Text())
		if !slices.Contains(versionsWithNoJarFile, version) {
			versions = append(versions, version)
		}
	})

	return versions, nil
}

func (f *Forge) DownloadURL(version string) (string, error) {
	versions, err := f.Versions()
	if err != nil {
		return "", err
	}

	if !slices.Contains(versions, version) {
		return "", fmt.Errorf("the specified version was not found: %s", version)
	}

	versionURL := fmt.Sprintf("%s/index_%s.html", filesURL, version)
	response, err := helpers.GetRequest(versionURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	var url string

	link := doc.Find("div.link-boosted a").First()
	url, exists := link.Attr("href")
	if !exists {
		return "", fmt.Errorf("href attribute not found for: %s", version)
	}

	url = strings.Split(url, "&url=")[1]

	return url, nil
}
