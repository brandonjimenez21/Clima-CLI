package model

import "fmt"

// ForecastPoint represents a single data point in the weather forecast
type ForecastPoint struct {
	Time string  `json:"Time"`
	Temp float64 `json:"Temp"`
}

// Weather data structure
type Weather struct {
	City             string
	Temperature      float64
	FeelsLike        float64
	Humidity         int
	Description      string
	Condition        string // Main condition: Clear, Clouds, Rain, etc.
	WindSpeed        float64
	ThermalSensation string          // Description of the thermal feeling
	Forecast         []ForecastPoint // Next 24 hours
	IsNight          bool            // True if it is night at the location
	UVIndex          float64         // UV Index
	Sunrise          string          // Formatted sunrise time
	Sunset           string          // Formatted sunset time
	RainProb         int             // Probability of precipitation (percentage)
	IsGoldenHour     bool            // True if it's sunrise or sunset
}

// CalculateThermalSensation returns a human-friendly string based on FeelsLike temperature
func (w *Weather) CalculateThermalSensation() {
	temp := w.FeelsLike
	switch {
	case temp <= 0:
		w.ThermalSensation = "Gélido"
	case temp > 0 && temp <= 10:
		w.ThermalSensation = "Frío"
	case temp > 10 && temp <= 20:
		w.ThermalSensation = "Fresco"
	case temp > 20 && temp <= 28:
		w.ThermalSensation = "Agradable"
	case temp > 28 && temp <= 35:
		w.ThermalSensation = "Caluroso"
	default:
		w.ThermalSensation = "Extremadamente Caluroso"
	}
}

func (w *Weather) DisplayString() string {
	return fmt.Sprintf("Ciudad: %s\nTemperatura: %.1f°C\nSensación Térmica: %.1f°C (%s)\nHumedad: %d%%\nClima: %s\nViento: %.1f m/s",
		w.City, w.Temperature, w.FeelsLike, w.ThermalSensation, w.Humidity, w.Description, w.WindSpeed)
}
