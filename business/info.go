package business

import (
	"fmt"
	"io"
	"net/http"

	"urbanmedia/go-cli-boilerplate/config"
)

// FetchInfo gets information from the API
func FetchInfo() (string, error) {
	apiURL := config.AppConfig.ApiURL

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
