package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/dan-frohlich/battlestations/character"
)

type App struct {
	model *mainModel
}

func NewApp() *App {
	theme := huh.ThemeBase16()

	a := &App{
		model: &mainModel{
			focus:   menuPanel,
			usecase: mainView,
			menu:    newMenuModel(theme),
			detail:  newDetailModel(),
			status:  newStatusModel(theme),
			file:    FileModel{},
			manager: &character.Manager{},
		},
	}
	a.model.app = a
	return a
}

func (app *App) Start() {
	p := tea.NewProgram(app.model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		//TODO should we panic?
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
