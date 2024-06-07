package utils

import (
	"dev-docs-cli/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchWeatherForecast fetches the weather forecast from the given URL
func FetchWeatherForecast(url string, token string) ([]models.WeatherForecast, error) {
	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the token in the request header
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read and parse the response
	var forecasts []models.WeatherForecast
	err = json.NewDecoder(resp.Body).Decode(&forecasts)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return forecasts, nil
}
