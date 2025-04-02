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
	view    viewport.Model
	content string
}

func newDetailModel() detailModel {
	return detailModel{
		view: viewport.New(12, 12),
	}
}

func (m detailModel) SetContent(content string) detailModel {
	m.content = content
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
		m.view = viewport.New(msg.Width, msg.Height)
		m.view.SetContent(m.content)
	}
	return m, nil

}

// View implements tea.Model.
func (m detailModel) View() string {
	return m.view.View()
}
