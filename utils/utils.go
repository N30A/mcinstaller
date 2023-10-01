package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func DecodeJSON(reader io.Reader, target interface{}) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(target)
}

func GetRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("request responded with status code: %d", response.StatusCode)
	}

	return response, nil
}

func DownloadFile(url, output string) error {

	response, err := GetRequest(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(output)
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
