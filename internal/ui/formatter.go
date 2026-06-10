package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/username/clima-cli/internal/model"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			MarginBottom(1)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2)
)

func FormatWeather(w *model.Weather) string {
	title := titleStyle.Render(fmt.Sprintf(" CLIMA EN %s ", strings.ToUpper(w.City)))

	content := strings.Builder{}
	content.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("🌡️  Temperatura:"), valueStyle.Render(fmt.Sprintf("%.1f°C", w.Temperature))))
	content.WriteString(fmt.Sprintf("%s %s (%s)\n", labelStyle.Render("🔥 Sensación:"), valueStyle.Render(fmt.Sprintf("%.1f°C", w.FeelsLike)), w.ThermalSensation))
	content.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("💧 Humedad:"), valueStyle.Render(fmt.Sprintf("%d%%", w.Humidity))))
	content.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("☁️  Estado:"), valueStyle.Render(strings.Title(w.Description))))
	content.WriteString(fmt.Sprintf("%s %s\n", labelStyle.Render("💨 Viento:"), valueStyle.Render(fmt.Sprintf("%.1f m/s", w.WindSpeed))))

	return title + "\n" + boxStyle.Render(content.String())
}

func FormatError(err error) string {
	errorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true).
		Padding(0, 1).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("#FF0000"))

	return errorStyle.Render(fmt.Sprintf("ERROR: %s", err.Error()))
}
