package models

type WeatherForecast struct {
	Date         string `json:"date"`
	TemperatureC int    `json:"temperatureC"`
	TemperatureF int    `json:"temperatureF"`
	Summary      string `json:"summary"`
}
