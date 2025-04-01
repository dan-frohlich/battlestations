package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// assert interface compliance
var _ tea.Model = &menuModel{}

func newMenuModel() menuModel {
	return menuModel{
		view: viewport.New(12, 12),
	}
}

func (m menuModel) SetContent(content string) menuModel {
	m.content = content
	m.view.SetContent(content)
	return m
}

func (m menuModel) DesiredWidth() (w int) {
	for _, s := range strings.Split(m.content, "\n") {
		i := len(s)
		if i > w {
			w = i
		}
	}
	return w + 2
}

type menuModel struct {
	view    viewport.Model
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
			// m.content += "\n" + msg.String()
			return m, tea.WindowSize()
		}
	// case tea.WindowSizeMsg:
	// 	m.view = viewport.New(12, msg.Height-6)
	// 	m.view.SetContent(m.content)
	case ContentMsg:
		m.content = string(msg)
		m.view.SetContent(m.content)
	case SizeMessage:
		m.view = viewport.New(msg.Width, msg.Height)
		m.view.SetContent(m.content)
	}
	return m, nil
}

// View implements tea.Model.
func (m menuModel) View() string {
	return m.view.View()
}
