package model

import "fmt"

// Weather data structure
type Weather struct {
	City             string
	Temperature      float64
	FeelsLike        float64
	Humidity         int
	Description      string
	WindSpeed        float64
	ThermalSensation string // Description of the thermal feeling
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
