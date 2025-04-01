package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	model *mainModel
}

func NewApp() *App {
	a := &App{
		model: &mainModel{
			focus:   menuPanel,
			usecase: mainView,
			menu: menuModel{
				content: "MY MENU",
				view:    viewport.New(12, 12),
			},
			detail: detailModel{},
			status: statusModel{},
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
