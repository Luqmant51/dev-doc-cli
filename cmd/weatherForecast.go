package cmd

import (
	"dev-docs-cli/pkg/models"
	"dev-docs-cli/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Fetch the weather forecast",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")

		// Get the executable path
		exePath, err := os.Executable()
		if err != nil {
			log.Fatalf("Failed to get executable path: %v", err)
		}

		// Construct the token file path
		rootPath := filepath.Dir(exePath)
		tokenPath := filepath.Join(rootPath, "devdocs_token")

		// Read the token from the file
		user, err := utils.ReadTokenFromFile(tokenPath)
		if err != nil {
			log.Fatalf("Failed to read token from file: %v", err)
		}

		// Check if the token is expired
		expired, err := utils.IsTokenExpired(user.AccessToken)
		if err != nil {
			log.Fatalf("Failed to check if token is expired: %v", err)
		}
		if expired {
			log.Fatal("Token has expired")
		}

		// Create a new request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Failed to create request: %v", err)
		}

		// Set the token in the request header
		req.Header.Set("Authorization", "Bearer "+user.AccessToken)

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Read and parse the response
		var forecasts []models.WeatherForecast
		err = json.NewDecoder(resp.Body).Decode(&forecasts)
		if err != nil {
			log.Fatalf("Failed to parse response: %v", err)
		}

		// Print the result
		for _, forecast := range forecasts {
			fmt.Printf("Date: %s\nTemperature (C): %d\nTemperature (F): %d\nSummary: %s\n\n",
				forecast.Date, forecast.TemperatureC, forecast.TemperatureF, forecast.Summary)
		}
	},
}

func init() {
	weatherCmd.Flags().String("url", "https://localhost:44319/WeatherForecast", "API URL to hit")
	rootCmd.AddCommand(weatherCmd)
}
