package tui

import tea "github.com/charmbracelet/bubbletea"

// assert interface compliance
var _ tea.Model = &statusModel{}

type statusModel struct {
}

// Init implements tea.Model.
func (d *statusModel) Init() tea.Cmd {
	panic("unimplemented")
}

// Update implements tea.Model.
func (d *statusModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (d *statusModel) View() string {
	panic("unimplemented")
}
