package tui

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/dan-frohlich/battlestations/character"
)

type App struct {
	theme   huh.Theme
	main    *mainModel
	overlay *olay
}

// Init implements tea.Model.
func (app *App) Init() tea.Cmd {
	return tea.Batch(app.main.Init(), app.overlay.Init())
}

// Update implements tea.Model.
func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logWithTS(time.Now(), fmt.Sprintf("called app.update([%[1]T][%[1]s])", msg))
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "f9": //quit
			return app, tea.Quit
		}
	}

	l, c := app.overlay.Update(msg)
	if ll, ok := l.(*olay); ok {
		app.overlay = ll
	}
	cmds = append(cmds, c)
	return app, tea.Batch(cmds...)
}

// View implements tea.Model.
func (app *App) View() string {
	return app.overlay.View()
}

var _ tea.Model = &App{}

func NewApp() *App {
	theme := huh.ThemeBase16()
	main := &mainModel{
		focus:   menuPanel,
		usecase: mainView,
		menu:    newMenuModel(theme),
		detail:  newDetailModel(),
		status:  newStatusModel(theme),
		file:    FileModel{},
		manager: &character.Manager{},
	}

	a := &App{
		overlay: newOverlay(main, theme),
		main:    main,
		theme:   *theme,
	}
	return a
}

func (app *App) Start() {
	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		//TODO should we panic?
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
