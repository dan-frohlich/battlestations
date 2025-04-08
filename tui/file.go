package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dan-frohlich/battlestations/character/model"
)

var _ tea.Model = FileModel{}

type FileModel struct {
}

// Init implements tea.Model.
func (f FileModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (f FileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case newCharFileMsg:
		cmds = append(cmds, newCmd(characterLoadedMsg{}), makeStatusCmd(infoLevel, "created character"))
		name := "[New Character]"
		cmds = append(cmds, newCmd(popupMsg{title: "Loaded Character", message: name}))
	case loadCharFileMsg:
		//TOTO refactor loading char as a tea.Cmd
		b, e := os.ReadFile(string(msg))
		if e != nil {
			return f, tea.Batch(
				makeStatusCmd(errorLevel, fmt.Sprintf("error loading character: %s", e)),
				makeUsecaseTransition(mainView))
		}

		c, e := model.LoadCharacter(b)
		if e != nil {
			return f, tea.Batch(
				makeStatusCmd(errorLevel, fmt.Sprintf("error parsing character: %s", e)),
				makeUsecaseTransition(mainView))
		}

		issues, errs := model.NewCharacterValidator().ValidateAll(c)
		for _, e := range errs {
			cmds = append(cmds, makeStatusCmd(errorLevel, e.Error()))
		}
		for _, issue := range issues {
			cmds = append(cmds, makeStatusCmd(warnLevel, issue.String()))
		}
		cmds = append(cmds, newCmd(characterLoadedMsg{c: c, sourceFile: string(msg)}), makeStatusCmd(infoLevel, "loaded character "+c.Name))
		name := c.Name
		if c.Name == "" {
			name = "[New Character]"
		}
		cmds = append(cmds, newCmd(popupMsg{title: "Loaded Character", message: name}))

	}

	return f, tea.Batch(cmds...)
}

// View implements tea.Model.
func (f FileModel) View() string {
	return "" // invisible component
}
