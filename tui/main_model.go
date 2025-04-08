package tui

import (
	"fmt"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dan-frohlich/battlestations/character"
	"github.com/dan-frohlich/battlestations/character/model"
)

// assert interface compliance
var _ tea.Model = &mainModel{}

type mainModel struct {
	focus   panel
	usecase usecaseView
	menu    menuModel
	detail  detailModel
	status  statusModel
	file    FileModel
	manager *character.Manager
}

var onStart = sync.Once{}

func (m mainModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.menu.Init())
	cmds = append(cmds, m.detail.Init())
	cmds = append(cmds, m.status.Init())
	onStart.Do(func() {
		cmds = append(cmds,
			//  makeUsecaseTransition(mainView),
			makeUsecaseTransition(mainHelpView))
	})
	return tea.Batch(cmds...)
}

func (m mainModel) charNeedsBasics() bool {
	c := m.manager.GetCharacter()
	return c.Name == "" ||
		c.Rank == 0 ||
		c.StartingSkillSet == "" ||
		c.Profession == "" ||
		c.Species.Name == "" ||
		len(c.Gear) < 1 ||
		len(c.SpecialAbilities) < 1
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// log(fmt.Sprintf("[%s] - called main.update([%[2]T][%[2]s])", time.Now().Format(time.RFC3339Nano), msg))
	// eventLog(msg)
	var (
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "f9":
			return m, tea.Quit
		case "tab":
			return m, makeFocusCmd(m.focus.next())
		case "shift+tab":
			return m, makeFocusCmd(m.focus.prev())
			// default:
			// 	return m, makeStatusCmd(debugLevel, "keypress: "+msg.String())
		}
	case displayHelpMsg:
		l, c := m.detail.Update(msg)
		if ll, ok := l.(detailModel); ok {
			m.detail = ll
		}
		return m, c
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
	case characterLoadedMsg:
		m.manager.SetCharacter(&msg.c)
		m.manager.SourceFile = msg.sourceFile
		l, c := m.detail.Update(msg)
		if ll, ok := l.(detailModel); ok {
			m.detail = ll
		}
		cmds = append(cmds, c)
		l, c = m.menu.Update(msg)
		if ll, ok := l.(menuModel); ok {
			m.menu = ll
		}
		cmds = append(cmds, c)
		v := manageCharView
		if m.charNeedsBasics() {
			v = newCharView
		}
		cmds = append(cmds, makeUsecaseTransition(v))
		return m, tea.Batch(cmds...)
	case characterModifiedMsg:
		l, c := m.detail.Update(msg)
		if ll, ok := l.(detailModel); ok {
			m.detail = ll
		}
		cmds = append(cmds, c)
		l, c = m.menu.Update(msg)
		if ll, ok := l.(menuModel); ok {
			m.menu = ll
		}
		cmds = append(cmds, c)
		v := manageCharView
		if m.charNeedsBasics() {
			v = newCharView
		}
		cmds = append(cmds, makeUsecaseTransition(v))
		return m, tea.Batch(cmds...)
	case selectedStatsMsg:
		m.manager.Modified = true
		c := m.manager.GetCharacter()
		c.StartingSkillSet = model.SkillSet(msg.selectedStatArray.String())
		lvls := c.StartingSkillSet.AsLevels()
		for i, skill := range msg.assignedSkills {
			switch skill {
			case skillAthletics:
				if c.Athletics == 0 {
					c.Athletics = lvls[i]
				}
			case skillCombat:
				if c.Combat == 0 {
					c.Combat = lvls[i]
				}
			case skillEngineering:
				if c.Engineering == 0 {
					c.Engineering = lvls[i]
				}
			case skillPilot:
				if c.Pilot == 0 {
					c.Pilot = lvls[i]
				}
			case skillScience:
				if c.Science == 0 {
					c.Science = lvls[i]
				}
			case optSkillDiplomacy:
				if c.Diplomacy == 0 {
					c.Diplomacy = model.OptionalSkillLevel(lvls[i])
				}
			case optSkillPsionics:
				if c.Psionics == 0 {
					c.Psionics = model.OptionalSkillLevel(lvls[i])
				}
			case optSkillSanity:
				if c.Sanity == 0 {
					c.Sanity = model.OptionalSkillLevel(lvls[i])
				}
			}
		}
		cmds = append(cmds, newCmd(characterModifiedMsg{c: c}))
	case panel:
		m.focus = msg
		return m, nil
	case usecaseView:
		prevView := m.usecase
		m.usecase = msg
		//we need to transition from m.usecase to msg
		l, c := m.menu.Update(msg)
		cmds = append(cmds, c)
		if ll, ok := l.(menuModel); ok {
			m.menu = ll
		}
		switch msg {
		case newCharView, mainView, loadCharSubView, manageCharView, setStatsSubView:
			//handle elsewhere
			l, c := m.menu.Update(msg)
			cmds = append(cmds, c)
			if ll, ok := l.(menuModel); ok {
				m.menu = ll
			}
			l, c = m.detail.Update(msg)
			cmds = append(cmds, c)
			if ll, ok := l.(detailModel); ok {
				m.detail = ll
			}
		case mainHelpView:
			l, c := m.menu.Update(msg)
			cmds = append(cmds, c)
			if ll, ok := l.(menuModel); ok {
				m.menu = ll
			}
			helpMsg := displayHelpMsg(
				`Welcome to the Battlestations Character Manager
 * to change focus: 'tab' or 'shift+tab'
 * to scroll up in a viewport: 'up', '8', 'w'
 * to scroll down in a viewport: 'down', '2', 's'
 * to page up in a viewport: 'pageup'
 * to page down in a viewport: 'pagedown'
 * to scroll to top of viewport: 'home'
 * to scroll bottom of viewport: 'end'
 * to return to main menu: 'esc'
 * to select an item from a menu: 'enter', 'return'
 * to open a folder in a file browser: 'right'
 * to open the parent folder in a file browser: 'left'
`)
			l, c = m.detail.Update(helpMsg)
			cmds = append(cmds, c)
			if ll, ok := l.(detailModel); ok {
				m.detail = ll
			}
		default:
			cmds = append(cmds, makeStatusCmd(errorLevel, fmt.Sprintf("failed to transition from view %s to view %s", prevView, msg)))
			if prevView == msg {
				cmds = append(cmds, makeUsecaseTransition(mainView)) //go back!
				return m, tea.Batch(cmds...)
			}
			cmds = append(cmds, makeUsecaseTransition(prevView)) //go back!
			return m, tea.Batch(cmds...)
		}
		// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("transitioning from view %s to view %s", m.usecase, msg)))
		return m, tea.Batch(cmds...)
	case newMenuMsg:
		var cmds []tea.Cmd
		sm, c := m.menu.Update(msg)
		cmds = append(cmds, c)
		if mm, ok := sm.(menuModel); ok {
			m.menu = mm
		}
		return m, tea.Batch(cmds...)
	case statusMsg:
		sm, c := m.status.Update(msg)
		if mm, ok := sm.(statusModel); ok {
			m.status = mm
			return m, c
		}
	case tea.WindowSizeMsg:
		widthOfPaddingAndBorder := 4
		heightOfPaddingAndBorder := 2
		statusViewPortHeight := 4
		menuWidth := lipgloss.Width(m.menu.View())
		menuResize := SizeMsg{
			Width:  menuWidth,
			Height: msg.Height - statusViewPortHeight - 2*heightOfPaddingAndBorder - 1,
		}
		detailResize := SizeMsg{
			Width:  msg.Width - menuWidth - 2*widthOfPaddingAndBorder,
			Height: msg.Height - statusViewPortHeight - 2*heightOfPaddingAndBorder,
		}
		statusResize := SizeMsg{
			Width:  msg.Width - widthOfPaddingAndBorder,
			Height: statusViewPortHeight,
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
			// cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("unhandled message %T", msg)))
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

var (
	borderStyle        = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(0, 1, 0, 1)
	focusedBorderStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), true).Padding(0, 1, 0, 1)
)

func (m mainModel) View() string {
	// log(fmt.Sprintf("[%s] - called detail.view", time.Now().Format(time.RFC3339Nano)))
	var (
		menuStyle   = borderStyle
		detailStyle = borderStyle
		statusStyle = borderStyle
	)
	switch m.focus {
	case menuPanel:
		menuStyle = focusedBorderStyle
	case detailPanel:
		detailStyle = focusedBorderStyle
	case statusPanel:
		statusStyle = focusedBorderStyle
	}
	menu := menuStyle.Render(m.menu.View())
	detail := detailStyle.Render(m.detail.View())
	status := statusStyle.Render(m.status.View())

	top := lipgloss.JoinHorizontal(lipgloss.Top, menu, detail)
	all := lipgloss.JoinVertical(lipgloss.Left, top, status)
	return all
}
