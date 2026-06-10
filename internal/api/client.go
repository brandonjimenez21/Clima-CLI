package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/username/clima-cli/internal/model"
)

const baseURL = "https://api.openweathermap.org/data/2.5/weather"

type openWeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type Client struct {
	APIKey string
}

func NewClient() (*Client, error) {
	// Intentar cargar desde .env si existe
	_ = godotenv.Load()

	apiKey := os.Getenv("CLIMA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("la variable CLIMA_API_KEY no se encontró en el entorno ni en el archivo .env")
	}
	return &Client{APIKey: apiKey}, nil
}

func (c *Client) GetWeather(city string) (*model.Weather, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric&lang=es", baseURL, city, c.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error de red: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("API Key inválida")
		}
		return nil, fmt.Errorf("ciudad no encontrada o error de API")
	}

	var data openWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error al procesar datos: %w", err)
	}

	weather := &model.Weather{
		City:        data.Name,
		Temperature: data.Main.Temp,
		FeelsLike:   data.Main.FeelsLike,
		Humidity:    data.Main.Humidity,
		Description: data.Weather[0].Description,
		WindSpeed:   data.Wind.Speed,
	}

	weather.CalculateThermalSensation()

	return weather, nil
}
