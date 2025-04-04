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
	msg     NewMenuMessage
	usecase usecaseView
	c       model.Character
}

func newMenuModel() menuModel {
	return menuModel{}
}

func (m menuModel) DesiredWidth() (w int) {
	return 24
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
	case NewMenuMessage:
		m.msg = msg
		m.sel = huh.NewSelect[string]().Options(huh.NewOptions(msg.items(m.c)...)...).Title(msg.title)
		m.form = huh.NewForm(huh.NewGroup(m.sel)).WithShowHelp(true)
		// cmd = tea.WindowSize()
		cmds = append(cmds, cmd, m.form.Init())
	case SizeMsg:
		if m.form != nil {
			m.form = m.form.WithWidth(msg.Width).WithHeight(msg.Height)
		}
		// m.content = m.form.View()
	case usecaseView:
		m.usecase = msg
		switch msg {
		case mainView:
			// cmds = append(cmds, makeStatusCmd(infoLevel, "selected main menu"))
			cmds = append(cmds, newCmd(
				NewMenuMessage{
					title: "Main Menu",
					keys:  []any{"Create Character", "Load Character", "Quit"},
					options: map[string]tea.Cmd{
						"Create Character": newCmd(newCharFileMsg{}),
						"Load Character":   makeUsecaseTransition(loadCharSubView),
						"Quit":             tea.Quit},
				}))
			cmds = append(cmds, tea.WindowSize())
		case newCharView:
			cmds = append(cmds,
				newCmd(NewMenuMessage{
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
			m.file = huh.NewFilePicker().
				Picking(true).
				DirAllowed(false).
				CurrentDirectory(cd).
				Title("Load a Character File").
				Description("Select a .yaml character file").
				AllowedTypes([]string{".yaml", ".yml"})
			m.form = huh.NewForm(huh.NewGroup(m.file)).WithShowHelp(true)
			cmds = append(cmds, m.file.Init(), tea.WindowSize())
		case manageCharView:
			cmds = append(cmds, newCmd(NewMenuMessage{
				title: "Manage Character",
				keys: []any{
					"main menu",
					"preview",
					"print",
					"save",
					"aftermath",
					"purchase gear",
				},
				options: map[string]tea.Cmd{
					"main menu":     makeUsecaseTransition(mainView),
					"preview":       makeUsecaseTransition(charPreviewSubView),
					"print":         makeUsecaseTransition(printCharSubView),
					"save":          makeUsecaseTransition(saveCharSubView),
					"aftermath":     makeUsecaseTransition(missionAftermathView),
					"purchase gear": makeUsecaseTransition(purchaseGearSubView),
				}}))
			cmds = append(cmds, tea.WindowSize())
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
