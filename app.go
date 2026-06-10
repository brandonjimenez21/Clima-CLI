package main

import (
	"context"
	"fmt"

	"github.com/username/clima-cli/internal/api"
	"github.com/username/clima-cli/internal/model"
)

// App struct
type App struct {
	ctx    context.Context
	client *api.Client
}

// NewApp creates a new App struct
func NewApp() *App {
	client, _ := api.NewClient()
	return &App{
		client: client,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetWeather returns the weather data for a given city
func (a *App) GetWeather(city string) (*model.Weather, error) {
	if a.client == nil {
		client, err := api.NewClient()
		if err != nil {
			return nil, err
		}
		a.client = client
	}

	weather, err := a.client.GetWeather(city)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el clima: %v", err)
	}

	return weather, nil
}

// GetLocalCity returns the city name based on IP
func (a *App) GetLocalCity() (string, error) {
	if a.client == nil {
		client, err := api.NewClient()
		if err != nil {
			return "", err
		}
		a.client = client
	}
	return a.client.GetLocalCity()
}
