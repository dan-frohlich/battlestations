package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type App struct {
	model *mainModel
}

func NewApp() *App {
	theme := huh.ThemeCharm()

	a := &App{
		model: &mainModel{
			focus:   menuPanel,
			usecase: mainView,
			menu:    newMenuModel(),
			detail:  newDetailModel().SetContent("DETAILED VIEW"),
			status:  newStatusModel(theme).SetContent("STATUS MESSAGE VIEW"),
			file:    FileModel{},
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
