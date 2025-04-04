package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// assert interface compliance
var _ tea.Model = &statusModel{}

type statusModel struct {
	view    viewport.Model
	content string
	theme   *huh.Theme
}

func newStatusModel(theme *huh.Theme) statusModel {
	return statusModel{
		view:  viewport.New(12, 12),
		theme: theme,
	}
}

func (m statusModel) SetContent(content string) statusModel {
	m.content = content
	m.view.SetContent(content)
	return m
}

// Init implements tea.Model.
func (m statusModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case SizeMsg:
		m.view.Width = msg.Width
		m.view.Height = msg.Height
	case statusMsg:
		m = m.handleStatusUpdate(msg)
	case tea.KeyMsg:
		viewportControl(&m.view, msg)
	}
	return m, tea.Batch(cmds...)
}
func (m statusModel) handleStatusUpdate(sm statusMsg) statusModel {
	var s string
	switch sm.lvl {
	case debugLevel:
		s = debugStyle.Render(sm.String())
	case infoLevel:
		s = infoStyle.Render(sm.String())
	case warnLevel:
		s = warnStyle.Render(sm.String())
	case errorLevel:
		s = errorStyle.Render(sm.String())
	default:
		s = statusStyle.Render(sm.String())
	}
	log(fmt.Sprintf("[%s] - [%s] - %s", sm.ts.Format(time.RFC3339Nano), sm.caller, s))
	for _, s := range sm.callstack {
		log(s)
	}

	old := strings.Split(m.content, "\n")
	if len(old) > 0 && old[len(old)-1] == s {
		return m //squelch log spams
	}
	new := append(old, s)
	m.content = strings.Join(new, "\n")
	m.view.SetContent(m.content)
	_ = m.view.GotoBottom()
	return m
}

func log(msg string) {

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()
	// Write the string to the file
	_, _ = file.WriteString(msg + "\n")
}

// func eventLog(msg tea.Msg) {

// 	ts := time.Now().Format(time.RFC3339Nano)
// 	var logMsg string
// 	switch tp := msg.(type) {
// 	case panel, usecaseView, SizeMsg, ContentMsg, loadCharFileMsg:
// 		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[2]s\n", ts, msg)
// 	case tea.KeyMsg:
// 		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[2]s\n", ts, msg)
// 	case statusMsg:
// 		//skip
// 	case model.Character:
// 		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[3]s\n", ts, msg, tp.Name)
// 	case NewMenuMessage:
// 		logMsg = fmt.Sprintf("[%s] - [%[2]T] %[3]s\n", ts, msg, tp.title)
// 	default:
// 		logMsg = fmt.Sprintf("[%s] - [%[2]T] %#[2]v\n", ts, msg)
// 	}
// 	log(logMsg)

// }

var (
	gray = lipgloss.AdaptiveColor{Light: "#DDDDDD", Dark: "#444444"}
	// green    = lipgloss.AdaptiveColor{Dark: "#50fa7b"}
	// purple   = lipgloss.AdaptiveColor{Dark: "#bd93f9"}
	red    = lipgloss.AdaptiveColor{Dark: "#ff5555"}
	yellow = lipgloss.AdaptiveColor{Dark: "#f1fa8c"}
	// normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
	// indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	// cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	// fuchsia  = lipgloss.Color("#F780E2")
	// green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	// red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}

	// normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
	// indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	// cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
	// fuchsia  = lipgloss.Color("#F780E2")
	// green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	// red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}

	// background = lipgloss.AdaptiveColor{Dark: "#282a36"}
	// selection  = lipgloss.AdaptiveColor{Dark: "#44475a"}
	// foreground = lipgloss.AdaptiveColor{Dark: "#f8f8f2"}
	comment = lipgloss.AdaptiveColor{Dark: "#6272a4"}
	// green      = lipgloss.AdaptiveColor{Dark: "#50fa7b"}
	// purple     = lipgloss.AdaptiveColor{Dark: "#bd93f9"}
	// red        = lipgloss.AdaptiveColor{Dark: "#ff5555"}
	// yellow     = lipgloss.AdaptiveColor{Dark: "#f1fa8c"}

	statusStyle = lipgloss.NewStyle().Foreground(comment)
	debugStyle  = lipgloss.NewStyle().Foreground(gray)
	infoStyle   = lipgloss.NewStyle().Foreground(comment)
	warnStyle   = lipgloss.NewStyle().Foreground(yellow)
	errorStyle  = lipgloss.NewStyle().Foreground(red)
)

// View implements tea.Model.
func (m statusModel) View() string {
	return m.view.View()
}
