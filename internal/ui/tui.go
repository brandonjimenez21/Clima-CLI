package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/username/clima-cli/internal/api"
	"github.com/username/clima-cli/internal/model"
)

type state int

const (
	inputting state = iota
	loading
	displaying
	errored
)

type Model struct {
	state     state
	textInput textinput.Model
	weather   *model.Weather
	err       error
	client    *api.Client
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Introduce una ciudad..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	client, _ := api.NewClient()

	return Model{
		state:     inputting,
		textInput: ti,
		client:    client,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

type weatherMsg *model.Weather
type errMsg error

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.state == inputting {
				city := m.textInput.Value()
				if city == "" {
					return m, nil
				}
				m.state = loading
				return m, m.fetchWeather(city)
			}
			if m.state == displaying || m.state == errored {
				m.state = inputting
				m.textInput.SetValue("")
				return m, nil
			}
		}

	case weatherMsg:
		m.state = displaying
		m.weather = msg
		return m, nil

	case errMsg:
		m.state = errored
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) fetchWeather(city string) tea.Cmd {
	return func() tea.Msg {
		if m.client == nil {
			return errMsg(fmt.Errorf("API Key no configurada"))
		}
		w, err := m.client.GetWeather(city)
		if err != nil {
			return errMsg(err)
		}
		return weatherMsg(w)
	}
}

func (m Model) View() string {
	var s string

	switch m.state {
	case inputting:
		s = fmt.Sprintf(
			"¿De qué ciudad quieres saber el clima?\n\n%s\n\n%s",
			m.textInput.View(),
			"(esc para salir)",
		)
	case loading:
		s = "Consultando satélites... 🛰️"
	case displaying:
		s = FormatWeather(m.weather) + "\n\nPresiona Enter para buscar otra ciudad."
	case errored:
		s = FormatError(m.err) + "\n\nPresiona Enter para intentar de nuevo."
	}

	return lipgloss.NewStyle().Margin(1, 2).Render(s)
}
