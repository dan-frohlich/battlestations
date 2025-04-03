package tui

import (
	"fmt"
	"runtime"
	"strings"
	"time"

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

type statusMsg struct {
	lvl       statusLevel
	msg       string
	ts        time.Time
	caller    string
	callstack []string
}

func (sm statusMsg) String() string {
	switch sm.lvl {
	case debugLevel:
		return "[d] " + sm.msg
	case infoLevel:
		return "[i] " + sm.msg
	case warnLevel:
		return "[w] " + sm.msg
	case errorLevel:
		return "[e] " + sm.msg
	default:
		return sm.msg
	}
}

// type debugMsg string
// type infoMsg string
// type warnMsg string
// type errorMsg string

type loadCharFileMsg string

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
	return tea.Batch(newCmd(v), makeStatusCmd(debugLevel, "you selected "+v.String()))
}

func (nmm NewMenuMessage) items() (result []string) {
	return nmm.keys
}

func makeStatusCmd(l statusLevel, s string) tea.Cmd {
	var (
		caller string
	)
	i := 1
	if _, file, line, ok := runtime.Caller(i); ok {
		for strings.HasSuffix(file, "messages.go") && ok {
			i++
			_, file, line, ok = runtime.Caller(i)
		}
		r := []rune(file)
		j := strings.LastIndex(file, "battlestations/")
		file = string(r[j+len("battlestations/"):])
		// z := strings.Split(file, "/")
		// file = strings.Join(z[len(z)-2:], "/")
		caller = fmt.Sprintf("%s:%d", file, line)
	}
	return func() tea.Msg {
		switch l {
		case debugLevel:
			return statusMsg{ts: time.Now(), lvl: l, msg: "[d] " + s, caller: caller}
		case infoLevel:
			return statusMsg{ts: time.Now(), lvl: l, msg: "[i] " + s, caller: caller}
		case warnLevel:
			return statusMsg{ts: time.Now(), lvl: l, msg: "[w] " + s, caller: caller}
		case errorLevel:
			return statusMsg{ts: time.Now(), lvl: l, msg: "[e] " + s, caller: caller}
		default:
			return statusMsg{ts: time.Now(), lvl: l, msg: s, caller: caller}
		}
	}
}

func makeFocusCmd(p panel) tea.Cmd {
	return func() tea.Msg {
		return p
	}
}
