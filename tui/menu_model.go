package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/dan-frohlich/battlestations/character/model"
)

// assert interface compliance
var (
	_ tea.Model = &menuModel{}
)

type menuModel struct {
	// title   string
	form    *huh.Form
	sel     *huh.Select[string]
	file    *huh.FilePicker
	theme   *huh.Theme
	msg     newMenuMsg
	usecase usecaseView
	c       model.Character
	size    SizeMsg
}

func newMenuModel(theme *huh.Theme) menuModel {
	return menuModel{theme: theme}
}

// Init implements tea.Model.
func (m menuModel) Init() tea.Cmd {
	if m.form != nil {
		return m.form.Init()
	}
	if m.file != nil {
		return m.file.Init()
	}
	return nil
}

// Update implements tea.Model.
func (m menuModel) Update(msg tea.Msg) (_ tea.Model, cmd tea.Cmd) {
	var cmds []tea.Cmd
	var l tea.Model
	switch msg := msg.(type) {
	case model.Character:
		m.c = msg
	case newMenuMsg:
		m.msg = msg
		m.sel = huh.NewSelect[string]().Options(huh.NewOptions(msg.items(m.c)...)...).Title(msg.title)
		m.form = huh.NewForm(huh.NewGroup(m.sel)).WithShowHelp(true).WithTheme(m.theme)
		//select a min size, the form with push out from there
		m.form.WithWidth(10)
		m.sel.WithWidth(10)
		if m.size.Height > 0 {
			m.form.WithHeight(m.size.Height - 1)
			m.sel.WithHeight(m.size.Height - 1)
		}
		// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("make form %s with sz %s", msg.title, m.size)))
		// cmd = tea.WindowSize()
		cmds = append(cmds, cmd, m.form.Init(), tea.WindowSize(), newCmd(menuLoaded(msg.title)))
	case SizeMsg:
		if m.file != nil {
			msg.Height += 2
		}
		if m.form != nil {
			m.form = m.form.WithWidth(msg.Width).WithHeight(msg.Height)
			// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("set form %s with sz %s", m.msg.title, m.size)))
		}
		if m.sel != nil {
			m.sel.WithWidth(msg.Width).WithHeight(msg.Height - 1)
		}
		if m.file != nil {
			m.file.WithWidth(msg.Width).WithHeight(msg.Height - 1)
		}
		m.size = msg
	case usecaseView:
		m.usecase = msg
		switch msg {
		case mainView, mainHelpView:
			// cmds = append(cmds, makeStatusCmd(infoLevel, "selected main menu"))
			cmds = append(cmds, newCmd(
				newMenuMsg{
					title: "Main Menu",
					keys:  []any{"Create Character", "Load Character", "Help", "Quit"},
					options: map[string]tea.Cmd{
						"Create Character": newCmd(newCharFileMsg{}),
						"Load Character":   makeUsecaseTransition(loadCharSubView),
						"Help":             makeUsecaseTransition(mainHelpView),
						"Quit":             tea.Quit},
				}))
			cmds = append(cmds, tea.WindowSize())
		case newCharView:
			cmds = append(cmds,
				newCmd(newMenuMsg{
					title: "Create Character",
					keys: []any{
						"main menu",
						menuItem{display: "set stats", visible: func(c model.Character) bool { return m.c.StartingSkillSet == "" }},
						menuItem{display: "set species", visible: func(c model.Character) bool { return m.c.Species.Name == "" }},
						menuItem{display: "set name", visible: func(c model.Character) bool { return m.c.Name == "" }},
						menuItem{display: "set profession", visible: func(c model.Character) bool { return m.c.Profession == "" }},
						menuItem{display: "set basic gear", visible: func(c model.Character) bool { return len(m.c.Gear) == 0 }},
					},
					options: map[string]tea.Cmd{
						"main menu":      makeUsecaseTransition(mainView),
						"set stats":      makeUsecaseTransition(setStatsSubView),
						"set species":    makeUsecaseTransition(setSpeciesSubView),
						"set name":       makeUsecaseTransition(setNameSubView),
						"set profession": makeUsecaseTransition(setProfessionSubView),
						"set basic gear": makeUsecaseTransition(setBasicGearSubView),
					}}))
			cmds = append(cmds, tea.WindowSize())
		case loadCharSubView:
			cd, _ := os.Getwd()
			const title = "Load a Character File"
			m.file = huh.NewFilePicker().
				Picking(true).
				DirAllowed(false).
				CurrentDirectory(cd).
				Title(title).
				Description("Select a .yaml character file").
				AllowedTypes([]string{".yaml", ".yml"})
			//select a min size, the form with push out from there
			m.file.WithWidth(10)
			m.file.WithHeight(10)
			m.form = huh.NewForm(huh.NewGroup(m.file)).WithShowHelp(true).WithTheme(m.theme)
			cmds = append(cmds, m.file.Init(), tea.WindowSize(), newCmd(menuLoaded(title)))
		case manageCharView:
			cmds = append(cmds, newCmd(newMenuMsg{
				title: "Manage Character",
				keys: []any{
					"main menu",
					"preview",
					"export to pdf",
					"save",
					"aftermath",
					"purchase gear",
				},
				options: map[string]tea.Cmd{
					"main menu":     makeUsecaseTransition(mainView),
					"preview":       makeUsecaseTransition(charPreviewSubView),
					"export to pdf": makeUsecaseTransition(exportCharToPDFSubView),
					"save":          makeUsecaseTransition(saveCharSubView),
					"aftermath":     makeUsecaseTransition(missionAftermathView),
					"purchase gear": makeUsecaseTransition(purchaseGearSubView),
				}}))
			cmds = append(cmds, tea.WindowSize())
		case exportCharToPDFSubView:
			//TODO use manager to print (export) the character
			//* use the manager to export the character to pdf

			// OPEN QUESTIONS
			//* Q: ask the user to name the export file or just use a default?
			//* A: TODO
			//* Q: ask the user to select the target folder or just use a default?
			//* A: TODO
			//* Q: ask the user if we should open the pdf with the default app or just do it by default?
			//* A: TODO
		}
		// case filepicker.MsgFileChosen:
	default:
		if m.form != nil {
			l, cmd = m.form.Update(msg)
			if lm, ok := l.(*huh.Form); ok {
				m.form = lm
			}
			cmds = append(cmds, cmd)
		}
	}
	if m.form != nil && m.form.State == huh.StateCompleted {
		switch m.usecase {
		case loadCharSubView:
			cmds = append(cmds, newCmd(loadCharFileMsg(fmt.Sprintf("%s", m.file.GetValue()))))
			m.form = nil
			m.file = nil
		default:
			var sv string
			m.sel.Value(&sv)
			// Form is completed, process the results
			// and potentially switch to a different view
			if actionCmd, ok := m.msg.options[sv]; ok && actionCmd != nil {
				cmds = append(cmds, actionCmd, tea.WindowSize())
			}
			m.form = nil
			m.sel = nil
		}
	}
	return m, tea.Batch(cmds...)
}

// View implements tea.Model.
func (m menuModel) View() string {
	switch m.usecase {
	case loadCharSubView:
		if m.file != nil {
			if m.file.GetValue() != "" {
				return fmt.Sprintf("Selected: %s", m.file.GetValue())
			}
			return m.file.View()
		}
	default:
		if m.form != nil {
			if m.form.State == huh.StateCompleted {
				if m.sel != nil {
					var sv string
					m.sel.Value(&sv)
					return sv
				}
			}
			return m.form.View()
		}
	}
	return "loading..."
}
