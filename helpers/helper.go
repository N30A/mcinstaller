package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
