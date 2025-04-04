package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dan-frohlich/battlestations/character/model"
	"gopkg.in/yaml.v2"
)

// assert interface compliance
var _ tea.Model = &detailModel{}

type detailModel struct {
	view viewport.Model
}

func newDetailModel() detailModel {
	return detailModel{
		view: viewport.New(12, 12),
	}
}

func (m detailModel) SetContent(content string) detailModel {
	m.view.SetContent(content)
	return m
}

// Init implements tea.Model.
func (m detailModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m detailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case model.Character:
		out, err := yaml.Marshal(msg)
		if err != nil {
			return m, makeStatusCmd(errorLevel, "failed to marshal char: "+msg.Name)
		}
		m.view.SetContent(string(out))
	case SizeMsg:
		m.view.Width = msg.Width
		m.view.Height = msg.Height
	case tea.KeyMsg:
		viewportControl(&m.view, msg)
	}
	return m, nil

}

func viewportControl(view *viewport.Model, msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// what key was pressed?
		switch msg.String() {
		case "up", "8", "w":
			_ = view.LineUp(1)
		case "down", "2", "s":
			_ = view.LineDown(1)
		case "ctrl+up", "ctrl+8", "ctrl+w":
			_ = view.HalfViewUp()
		case "ctrl+down", "ctrl+2", "ctrl+s":
			_ = view.HalfViewDown()
		case "pgup":
			_ = view.ViewUp()
		case "pgdown":
			_ = view.ViewDown()
		case "home":
			_ = view.GotoTop()
		case "end":
			_ = view.GotoBottom()
			// default:
			// 	eventLog(msg)
		}
	}
}

// View implements tea.Model.
func (m detailModel) View() string {
	return m.view.View()
}
