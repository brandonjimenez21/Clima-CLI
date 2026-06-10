package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/username/clima-cli/internal/model"
)

const (
	baseURL      = "https://api.openweathermap.org/data/2.5/weather"
	forecastURL  = "https://api.openweathermap.org/data/2.5/forecast"
	pollutionURL = "https://api.openweathermap.org/data/2.5/air_pollution"
	geoIPURL     = "https://ipapi.co/json/"
)

type pollutionResponse struct {
	List []struct {
		Main struct {
			Aqi int `json:"aqi"`
		} `json:"main"`
	} `json:"list"`
}

type openWeatherResponse struct {
	Name  string `json:"name"`
	Dt    int64  `json:"dt"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type forecastResponse struct {
	List []struct {
		Dt   int64   `json:"dt"`
		Pop  float64 `json:"pop"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	} `json:"list"`
}

type Client struct {
	APIKey string
}

func NewClient() (*Client, error) {
	_ = godotenv.Load()
	apiKey := os.Getenv("CLIMA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("la variable CLIMA_API_KEY no se encontró")
	}
	return &Client{APIKey: apiKey}, nil
}

func (c *Client) GetLocalCity() (string, error) {
	resp, err := http.Get(geoIPURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		City string `json:"city"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.City, nil
}

func (c *Client) GetWeather(city string) (*model.Weather, error) {
	// 1. Current Weather (Search by name to get coordinates)
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric&lang=es", baseURL, city, c.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error de red: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ciudad no encontrada")
	}

	var data openWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// 2. Forecast using Coordinates for higher precision as requested
	fUrl := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=metric&lang=es", forecastURL, data.Coord.Lat, data.Coord.Lon, c.APIKey)
	fResp, err := http.Get(fUrl)
	var forecastPoints []model.ForecastPoint
	rainProb := 0
	if err == nil && fResp.StatusCode == http.StatusOK {
		var fData forecastResponse
		if err := json.NewDecoder(fResp.Body).Decode(&fData); err == nil {
			if len(fData.List) > 0 {
				rainProb = int(fData.List[0].Pop * 100)
			}
			for i, item := range fData.List {
				if i >= 8 {
					break
				}
				t := time.Unix(item.Dt, 0)
				forecastPoints = append(forecastPoints, model.ForecastPoint{
					Time: t.Format("15:04"),
					Temp: item.Main.Temp,
				})
			}
		}
		fResp.Body.Close()
	}

	isNight := data.Dt < data.Sys.Sunrise || data.Dt > data.Sys.Sunset

	// Golden hour: within 60 minutes of sunrise or sunset
	const goldenHourLimit = 3600
	isGoldenHour := (data.Dt >= data.Sys.Sunrise-goldenHourLimit && data.Dt <= data.Sys.Sunrise+goldenHourLimit) ||
		(data.Dt >= data.Sys.Sunset-goldenHourLimit && data.Dt <= data.Sys.Sunset+goldenHourLimit)

	weather := &model.Weather{
		City:         data.Name,
		Temperature:  data.Main.Temp,
		FeelsLike:    data.Main.FeelsLike,
		Humidity:     data.Main.Humidity,
		Description:  data.Weather[0].Description,
		Condition:    data.Weather[0].Main,
		WindSpeed:    data.Wind.Speed,
		Forecast:     forecastPoints,
		IsNight:      isNight,
		IsGoldenHour: isGoldenHour,
		Sunrise:      time.Unix(data.Sys.Sunrise, 0).Format("15:04"),
		Sunset:       time.Unix(data.Sys.Sunset, 0).Format("15:04"),
		RainProb:     rainProb,
		UVIndex:      0,
	}

	// 3. Air Quality (AQI)
	pUrl := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", pollutionURL, data.Coord.Lat, data.Coord.Lon, c.APIKey)
	pResp, err := http.Get(pUrl)
	if err == nil && pResp.StatusCode == http.StatusOK {
		var pData pollutionResponse
		if err := json.NewDecoder(pResp.Body).Decode(&pData); err == nil && len(pData.List) > 0 {
			weather.AirQuality = pData.List[0].Main.Aqi
		}
		pResp.Body.Close()
	}

	weather.CalculateThermalSensation()

	return weather, nil
}
