package tui

import tea "github.com/charmbracelet/bubbletea"

// assert interface compliance
var _ tea.Model = &detailModel{}

type detailModel struct {
}

// Init implements tea.Model.
func (m detailModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m detailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// what key was pressed?
		switch msg.String() {

		//  exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}
	return m, nil

}

// View implements tea.Model.
func (m detailModel) View() string {
	return "my details"
}
