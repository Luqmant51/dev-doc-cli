package cmd

import (
	"dev-docs-cli/pkg/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Fetch the weather forecast",
	Run: func(cmd *cobra.Command, args []string) {
		baseurl := os.Getenv("API_URL")
		fmt.Print(baseurl)
		if baseurl == "" {
			log.Fatal("API_URL environment variable is not set")
		}

		// Check if the baseurl has the proper scheme
		if !strings.HasPrefix(baseurl, "http://") && !strings.HasPrefix(baseurl, "https://") {
			log.Fatal("API_URL must include the protocol scheme (http:// or https://)")
		}

		url := baseurl + "/WeatherForecast"

		// Debug statement to verify URL
		fmt.Printf("Requesting weather forecast from URL: %s\n", url)

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

		// Fetch the weather forecast
		resp, err := utils.FetchWeatherForecast(url, user.AccessToken)
		if err != nil {
			log.Fatalf("Failed to fetch weather forecast: %v", err)
		}

		// Print the result
		for _, forecast := range resp {
			fmt.Printf("Date: %s\nTemperature (C): %d\nTemperature (F): %d\nSummary: %s\n\n",
				forecast.Date, forecast.TemperatureC, forecast.TemperatureF, forecast.Summary)
		}
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
