package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/username/clima-cli/internal/api"
	"github.com/username/clima-cli/internal/ui"
)

var rootCmd = &cobra.Command{
	Use:   "clima-cli",
	Short: "Clima CLI es una herramienta para consultar el clima en tiempo real",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(ui.NewModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error al iniciar la TUI: %v", err)
			os.Exit(1)
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get [ciudad]",
	Short: "Obtiene el clima de una ciudad específica directamente",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := args[0]
		client, err := api.NewClient()
		if err != nil {
			fmt.Println(ui.FormatError(err))
			os.Exit(1)
		}
		weather, err := client.GetWeather(city)
		if err != nil {
			fmt.Println(ui.FormatError(err))
			os.Exit(1)
		}
		fmt.Println(ui.FormatWeather(weather))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
