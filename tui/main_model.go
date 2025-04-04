package tui

import (
	"fmt"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dan-frohlich/battlestations/character/model"
)

// assert interface compliance
var _ tea.Model = &mainModel{}

type mainModel struct {
	app     *App
	focus   panel
	usecase usecaseView
	menu    menuModel
	detail  detailModel
	status  statusModel
	file    FileModel
}

var onStart = sync.Once{}

func (m mainModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.menu.Init())
	cmds = append(cmds, m.detail.Init())
	cmds = append(cmds, m.status.Init())
	onStart.Do(func() {
		cmds = append(cmds, makeUsecaseTransition(mainView))
	})
	return tea.Batch(cmds...)
}

func (m mainModel) eventLog(msg tea.Msg) {

	ts := time.Now().Format(time.RFC3339Nano)
	var logMsg string
	switch tp := msg.(type) {
	case panel, usecaseView, SizeMsg, ContentMsg, loadCharFileMsg:
		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[2]s\n", ts, msg)
	case statusMsg:
		//skip
	case model.Character:
		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[3]s\n", ts, msg, tp.Name)
	case NewMenuMessage:
		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[3]s\n", ts, msg, tp.title)
	default:
		logMsg = fmt.Sprintf("[%s] - [%[2]T] %#[2]v\n", ts, msg)
	}
	log(logMsg)

}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.eventLog(msg)
	var (
		cmds []tea.Cmd
	)
	//handle global commands: exit / quite / tab
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		// what key was pressed?
		switch msg.String() {

		//  exit the program.
		case "f9":
			return m, tea.Quit
		case "tab":
			return m, makeFocusCmd(m.focus.next())
		case "shift+tab":
			return m, makeFocusCmd(m.focus.prev())
			// default:
			// 	return m, makeStatusCmd(debugLevel, "keypress: "+msg.String())
		}
	case newCharFileMsg:
		l, c := m.file.Update(msg)
		if ll, ok := l.(FileModel); ok {
			m.file = ll
		}
		return m, c
	case loadCharFileMsg:
		l, c := m.file.Update(msg)
		if ll, ok := l.(FileModel); ok {
			m.file = ll
		}
		return m, c
	case model.Character:
		l, c := m.detail.Update(msg)
		if ll, ok := l.(detailModel); ok {
			m.detail = ll
		}
		v := manageCharView
		if msg.Name == "" ||
			msg.Rank == 0 ||
			msg.StartingSkillSet == "" ||
			msg.Profession == "" ||
			msg.Species.Name == "" ||
			len(msg.Gear) < 1 ||
			len(msg.SpecialAbilities) < 1 {
			v = newCharView
		}
		return m, tea.Batch(c, makeUsecaseTransition(v))
	case panel:
		m.focus = msg
		return m, nil
	case usecaseView:
		m.usecase = msg
		//we need to transition from m.usecase to msg
		l, c := m.menu.Update(msg)
		cmds = append(cmds, c)
		if ll, ok := l.(menuModel); ok {
			m.menu = ll
		}
		switch msg {
		case newCharView:
		case mainView:
		case loadCharSubView:
		case manageCharView:
		default:
			cmds = append(cmds, makeStatusCmd(errorLevel, fmt.Sprintf("failed to transition from view %s to view %s", m.usecase, msg)))
			cmds = append(cmds, makeUsecaseTransition(m.usecase)) //go back!
			return m, tea.Batch(cmds...)
		}
		// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("transitioning from view %s to view %s", m.usecase, msg)))
		return m, tea.Batch(cmds...)
	case NewMenuMessage:
		sm, c := m.menu.Update(msg)
		if mm, ok := sm.(menuModel); ok {
			m.menu = mm
			return m, c
		}
	case statusMsg:
		sm, c := m.status.Update(msg)
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			return m, c
		}
	case tea.WindowSizeMsg:
		menuWidth := m.menu.DesiredWidth()
		if maximizeMenu(m.usecase) {
			menuWidth = msg.Width - 2
		}
		menuResize := SizeMsg{
			Width:  menuWidth,
			Height: msg.Height - 9,
		}
		detailResize := SizeMsg{
			Width:  msg.Width - menuWidth - 23,
			Height: msg.Height - 8,
		}
		statusResize := SizeMsg{
			Width:  msg.Width - 4,
			Height: 4,
		}
		sm, c := m.menu.Update(menuResize)
		if mm, ok := sm.(menuModel); ok {
			m.menu = mm
			cmds = append(cmds, c)
		}
		sm, c = m.detail.Update(detailResize)
		if mm, ok := sm.(detailModel); ok {
			m.detail = mm
			cmds = append(cmds, c)
		}
		sm, c = m.status.Update(statusResize)
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			cmds = append(cmds, c)
		}
		// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("win size: %s", SizeMsg{Height: msg.Height, Width: msg.Width})))
		// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("set sizes: m:%s d:%s s%s", menuResize, detailResize, statusResize)))
		return m, tea.Batch(cmds...)

	default:
		t := fmt.Sprintf("%T", msg)
		switch t {
		case "huh.nextFieldMsg", "huh.nextGroupMsg", "tea.windowSizeMsg", "filepicker.readDirMsg":
		default:
			cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("unhandled message %T", msg)))
		}
	}

	switch m.focus {
	case menuPanel:
		sm, c := m.menu.Update(msg)
		if mm, ok := sm.(menuModel); ok {
			m.menu = mm
			cmds = append(cmds, c)
		}
	case detailPanel:
		sm, c := m.detail.Update(msg)
		if mm, ok := sm.(detailModel); ok {
			m.detail = mm
			cmds = append(cmds, c)
		}
	case statusPanel:
		sm, c := m.status.Update(msg)
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			cmds = append(cmds, c)
		}
	}

	return m, tea.Batch(cmds...)
}

func maximizeMenu(v usecaseView) bool {
	return v == loadCharSubView
}

var (
	style      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(0, 1, 0, 1)
	focusStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), true).Padding(0, 1, 0, 1)
)

func (m mainModel) View() string {
	var (
		menuStyle   = style
		detailStyle = style
		statusStyle = style
	)
	switch m.focus {
	case menuPanel:
		menuStyle = focusStyle
	case detailPanel:
		detailStyle = focusStyle
	case statusPanel:
		statusStyle = focusStyle
	}
	menu := menuStyle.Render(m.menu.View())
	detail := detailStyle.Render(m.detail.View())
	status := statusStyle.Render(m.status.View())

	top := lipgloss.JoinHorizontal(lipgloss.Top, menu, detail)
	all := lipgloss.JoinVertical(lipgloss.Left, top, status)
	return all
}
