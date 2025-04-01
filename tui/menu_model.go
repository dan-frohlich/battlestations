package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// assert interface compliance
var _ tea.Model = &menuModel{}

type menuModel struct {
	content string
}

// Init implements tea.Model.
func (m menuModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// what key was pressed?
		switch msg.String() {

		//  exit the program.
		case "ctrl+c", "esc":
			return m, tea.Quit
		default:
			m.content += "\n" + msg.String()
			return m, tea.WindowSize()
		}
	case ContentMsg:
		m.content = string(msg)
		// log.Println(msg)
	}
	return m, nil
}

// View implements tea.Model.
func (m menuModel) View() string {
	return m.content
}
