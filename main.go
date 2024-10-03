package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type AppState struct {
	count   int
	configs ServerConfigs
}

// Initial state of the application
func (m AppState) Init() tea.Cmd {
	return nil
}

// Handles input (like key presses) and updates the model
func (m AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up":
			m.count++
		case "down":
			m.count--
		}
	}
	return m, nil
}

// Draw the UI based on the current state of the model
func (m AppState) View() string {
	return fmt.Sprintf("Count: %d\nPress up/down to change count, q to quit.", m.count, m.configs)
}

func main() {
	// Initialize the program with an initial model

	sc := ServerConfigs{}
	sc.loadFilenames("./config")

	p := tea.NewProgram(AppState{count: 0, configs: sc})

	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
