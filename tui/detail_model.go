package tui

import tea "github.com/charmbracelet/bubbletea"

// assert interface compliance
var _ tea.Model = &detailModel{}

type detailModel struct {
}

// Init implements tea.Model.
func (d *detailModel) Init() tea.Cmd {
	panic("unimplemented")
}

// Update implements tea.Model.
func (d *detailModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (d *detailModel) View() string {
	panic("unimplemented")
}
