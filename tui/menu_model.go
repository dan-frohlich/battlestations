package tui

import tea "github.com/charmbracelet/bubbletea"

// assert interface compliance
var _ tea.Model = &menuModel{}

type menuModel struct {
}

// Init implements tea.Model.
func (d *menuModel) Init() tea.Cmd {
	panic("unimplemented")
}

// Update implements tea.Model.
func (d *menuModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (d *menuModel) View() string {
	panic("unimplemented")
}
