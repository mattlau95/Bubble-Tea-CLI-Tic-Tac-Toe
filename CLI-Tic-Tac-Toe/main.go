package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the model from ui.go
	m := initialModel()

	// Create the Bubble Tea program
	p := tea.NewProgram(m, tea.WithAltScreen())

	// Run the program and handle potential errors
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Tic-Tac-Toe: %v", err)
		os.Exit(1)
	}
}