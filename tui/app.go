package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	model *mainModel
}

func NewApp() *App {
	a := &App{
		model: initialModel(),
	}
	a.model.app = a

	return a
}

func (app *App) Start() {
	p := tea.NewProgram(app.model)

	if _, err := p.Run(); err != nil {
		//TODO should we panic?
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

func initialModel() *mainModel {
	return &mainModel{}
}
