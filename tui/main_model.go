package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// assert interface compliance
var _ tea.Model = &mainModel{}

type mainModel struct {
	app *App
}

func (m *mainModel) Init() tea.Cmd {
	return nil
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// what key was pressed?
		switch msg.String() {

		//  exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":

		// The "down" and "j" keys move the cursor down
		case "down", "j":

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
		}
	}

	// Return the updated mainModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m *mainModel) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	return s
}
