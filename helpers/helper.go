package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("received status code: %d", response.StatusCode)
	}

	return response, nil
}

func FetchAndDecodeJSON(url string, data interface{}) error {
	response, err := GetRequest(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFile(url, filePath string) error {
	response, err := GetRequest(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
