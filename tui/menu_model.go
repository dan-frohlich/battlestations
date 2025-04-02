package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

// assert interface compliance
var (
	_ tea.Model = &menuModel{}
)

type menuModel struct {
	title   string
	content string
	form    *huh.Form
	sel     *huh.Select[string]
	msg     NewMenuMessage
}

func newMenuModel(msg NewMenuMessage) menuModel {

	s := huh.NewSelect[string]().Options(huh.NewOptions(msg.items()...)...).Title(msg.title)

	m := menuModel{
		title: msg.title,
		form:  huh.NewForm(huh.NewGroup(s)),
		sel:   s,
		msg:   msg,
	}
	// s.Value(&m.chosen)

	return m
}

func (m menuModel) DesiredWidth() (w int) {
	return 24
	// for _, s := range m.msg.keys {
	// 	i := len([]rune(s))
	// 	if i > w {
	// 		w = i
	// 	}
	// }
	// return w + 8
}

// Init implements tea.Model.
func (m menuModel) Init() tea.Cmd {
	return m.form.Init()
}

// Update implements tea.Model.
func (m menuModel) Update(msg tea.Msg) (_ tea.Model, cmd tea.Cmd) {
	var cmds []tea.Cmd
	var l tea.Model
	switch msg := msg.(type) {
	case NewMenuMessage:
		m.msg = msg
		m.sel = huh.NewSelect[string]().Options(huh.NewOptions(msg.items()...)...).Title(msg.title)
		m.form = huh.NewForm(huh.NewGroup(m.sel))
		m.content = m.form.View()
		cmd = tea.WindowSize()
		cmds = append(cmds, cmd, m.form.Init())
		// cmds = append(cmds, makeStatusCmd(infoLevel, fmt.Sprintf("new menu title: %s", msg.title)))
	case SizeMsg:
		m.form = m.form.WithWidth(msg.Width).WithHeight(msg.Height)
		m.content = m.form.View()
	default:
		l, cmd = m.form.Update(msg)
		if lm, ok := l.(*huh.Form); ok {
			m.form = lm
		}
		m.content = l.View()
		cmds = append(cmds, cmd)
	}
	if m.form.State == huh.StateCompleted {
		var sv string
		m.sel.Value(&sv)
		// Form is completed, process the results
		// and potentially switch to a different view
		if actionCmd, ok := m.msg.options[sv]; ok && actionCmd != nil {
			cmds = append(cmds, actionCmd, tea.WindowSize())
		}
		// m.form.State = huh.StateNormal
	}
	return m, tea.Batch(cmds...)
}

// View implements tea.Model.
func (m menuModel) View() string {
	if m.form.State == huh.StateCompleted {
		var sv string
		m.sel.Value(&sv)
		return sv
	}

	return m.form.View()
}
