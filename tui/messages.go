package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ContentMsg string

type SizeMsg struct {
	Width  int
	Height int
}

func (sm SizeMsg) String() string {
	return fmt.Sprintf("{%d,%d}", sm.Width, sm.Height)
}

type statusLevel int

const (
	_ statusLevel = iota
	debugLevel
	infoLevel
	warnLevel
	errorLevel
)

type statusMsg string
type debugMsg string
type infoMsg string
type warnMsg string
type errorMsg string

type NewMenuMessage struct {
	title   string
	options map[string]tea.Cmd
	keys    []string
}

func newCmd(message tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return message
	}
}

func makeUsecaseTransition(v usecaseView) tea.Cmd {
	return tea.Batch(
		makeStatusCmd(infoLevel, "you selected "+v.String()),
		func() tea.Msg {
			return v
		})
}

func (nmm NewMenuMessage) items() (result []string) {
	return nmm.keys
}

func makeStatusCmd(l statusLevel, s string) tea.Cmd {
	return func() tea.Msg {
		switch l {
		case debugLevel:
			return debugMsg("[d] " + s)
		case infoLevel:
			return infoMsg("[i] " + s)
		case warnLevel:
			return warnMsg("[w] " + s)
		case errorLevel:
			return errorMsg("[e] " + s)
		default:
			return statusMsg(s)
		}
	}
}

func makeFocusCmd(p panel) tea.Cmd {
	return func() tea.Msg {
		return p
	}
}
