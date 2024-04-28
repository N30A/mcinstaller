package helpers

import (
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
