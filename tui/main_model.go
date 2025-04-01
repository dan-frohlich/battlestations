package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// assert interface compliance
var _ tea.Model = &mainModel{}

type mainModel struct {
	app     *App
	focus   panel
	usecase usecaseView
	menu    menuModel
	detail  detailModel
	status  statusModel
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)
	//handle global commands: exit / quite / tab
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// what key was pressed?
		switch msg.String() {

		//  exit the program.
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab":
			m.focus = m.focus.next()
			return m, nil
		case "shift+tab":
			m.focus = m.focus.prev()
			return m, nil

			// // The "up" and "k" keys move the cursor up
			// case "up", "k":

			// // The "down" and "j" keys move the cursor down
			// case "down", "j":

			// // The "enter" key and the spacebar (a literal space) toggle
			// // the selected state for the item that the cursor is pointing at.
			// case "enter", " ":
		}
	case usecaseView:
		//we need to transition from m.usecase to msg
		sm, c := m.status.Update(fmt.Sprintf("transitioning from view %s to view %s", m.usecase, msg))
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			cmds = append(cmds, c)
		}
	case tea.WindowSizeMsg:
		menuWidth := m.menu.DesiredWidth()
		menuResize := SizeMessage{
			Width:  menuWidth,
			Height: msg.Height - 8,
		}
		detailResize := SizeMessage{
			Width:  msg.Width - menuWidth - 8,
			Height: msg.Height - 8,
		}
		statusResize := SizeMessage{
			Width:  msg.Width - 4,
			Height: 4,
		}
		sm, c := m.menu.Update(menuResize)
		if mm, ok := sm.(menuModel); ok {
			m.menu = mm
			cmds = append(cmds, c)
		}
		sm, c = m.detail.Update(detailResize)
		if mm, ok := sm.(detailModel); ok {
			m.detail = mm
			cmds = append(cmds, c)
		}
		sm, c = m.status.Update(statusResize)
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			cmds = append(cmds, c)
		}
		return m, tea.Batch(cmds...)
	}
	switch m.focus {
	case menuPanel:
		sm, c := m.menu.Update(msg)
		if mm, ok := sm.(menuModel); ok {
			m.menu = mm
			cmds = append(cmds, c)
		}
	case detailPanel:
		sm, c := m.detail.Update(msg)
		if mm, ok := sm.(detailModel); ok {
			m.detail = mm
			cmds = append(cmds, c)
		}
	case statusPanel:
		sm, c := m.status.Update(msg)
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			cmds = append(cmds, c)
		}
	}

	return m, tea.Batch(cmds...)
}

var (
	style      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(0, 1, 0, 1)
	focusStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), true).Padding(0, 1, 0, 1)
)

func (m mainModel) View() string {
	var (
		menuStyle   = style
		detailStyle = style
		statusStyle = style
	)
	switch m.focus {
	case menuPanel:
		menuStyle = focusStyle
	case detailPanel:
		detailStyle = focusStyle
	case statusPanel:
		statusStyle = focusStyle
	}
	menu := menuStyle.Render(m.menu.View())
	detail := detailStyle.Render(m.detail.View())
	status := statusStyle.Render(m.status.View())

	top := lipgloss.JoinHorizontal(lipgloss.Top, menu, detail)
	all := lipgloss.JoinVertical(lipgloss.Left, top, status)
	return all
}
