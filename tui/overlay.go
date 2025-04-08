package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

var _ tea.Model = &olay{}

type olay struct {
	displayOverlay bool
	usecase        usecaseView
	winHeight      int
	fg             *popup
	bg             tea.Model
	overlay        *overlay.Model
	theme          *huh.Theme
}

func newOverlay(bg tea.Model, theme *huh.Theme) *olay {
	fg := &popup{
		title:   "Popup Title",
		message: "popup message to display",
		theme:   theme,
	}
	return &olay{
		fg: fg,
		bg: bg,
		overlay: overlay.New(
			fg,
			bg,
			overlay.Center,
			overlay.Center,
			0,
			0,
		),
		displayOverlay: false,
		theme:          theme,
	}
}

// Init implements tea.Model.
func (o *olay) Init() tea.Cmd {
	return tea.Batch(o.overlay.Init(), o.fg.Init(), o.bg.Init())
}

// Update implements tea.Model.
func (o *olay) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logWithTS(time.Now(), fmt.Sprintf("called olay.update([%[1]T][%[1]s])", msg))
	var cmds []tea.Cmd
	//cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("overlay: [%T[1] - %[1]s", stringMsg(msg))))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// log(fmt.Sprintf("[%s] - called olay.update([%[2]T][%[2]s])", time.Now().Format(time.RFC3339Nano), msg))
		// cmds = append(cmds, makeStatusCmd(infoLevel, fmt.Sprintf("overlay received: [%T[1] - %[1]s", stringMsg(msg))))
		if o.displayOverlay {
			if o.fg.nested != nil {
				switch msg.String() {
				case "esc":
					o.displayOverlay = false
					return o, nil
				}
				// l, c := o.fg.Update(msg)
				// if ll, ok := l.(*popup); ok {
				// 	o.fg = ll
				// }
				// cmds = append(cmds, c)
			} else {
				// what key was pressed?
				switch msg.String() {
				case "esc", "enter":
					o.displayOverlay = false
				}
				return o, nil
			}
		}
	case tea.WindowSizeMsg:
		o.winHeight = msg.Height
	case popupMsg:
		o.displayOverlay = true
	case usecaseView:
		o.usecase = msg
		switch msg {
		case setStatsSubView:
			cmds = append(cmds, newCmd(popupMsg{nested: newSetStatsModel(o.theme)}))
		}
	case selectedStatsMsg:
		o.displayOverlay = false
	}

	l, cmd := o.fg.Update(msg)
	if ll, ok := l.(*popup); ok {
		o.fg = ll
	}
	cmds = append(cmds, cmd)

	l, cmd = o.overlay.Update(msg)
	if ll, ok := l.(*overlay.Model); ok {
		o.overlay = ll
	}
	cmds = append(cmds, cmd)

	o.bg, cmd = o.bg.Update(msg)
	cmds = append(cmds, cmd)

	return o, tea.Batch(cmds...)
}

// View implements tea.Model.
func (o *olay) View() string {
	// log(fmt.Sprintf("[%s] - %s", time.Now().Format(time.RFC3339Nano), "called olay.view"))
	if o.displayOverlay {
		if o.winHeight > 0 {
			popH := lipgloss.Height(o.fg.View())
			yOffset := (o.winHeight - popH) / 3
			o.overlay.YOffset = -1 * yOffset
		}
		o.overlay.Background = o.bg
		return o.overlay.View()
	}
	return o.bg.View()
}

var _ tea.Model = &popup{}

type hasForm interface {
	getForm() *huh.Form
	getResultCmd() tea.Cmd
}

type popup struct {
	title   string
	message string
	// help    string
	theme  *huh.Theme
	nested tea.Model
}

// Init implements tea.Model.
func (p *popup) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (p *popup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	logWithTS(time.Now(), fmt.Sprintf("called popup.update([%[1]T][%[1]s])", msg))
	switch msg := msg.(type) {
	case popupMsg:
		p.title = msg.title
		p.message = msg.message
		p.nested = msg.nested
		return p, tea.WindowSize()
	case tea.WindowSizeMsg:
		if p.nested != nil {
			p.nested.Update(tea.WindowSizeMsg{Height: msg.Height, Width: 80})
		}
	default:
		if p.nested != nil {
			var c tea.Cmd
			p.nested, c = p.nested.Update(msg)
			cmds = append(cmds, c)
		}
	}
	if p.nested != nil {
		if fm, ok := p.nested.(hasForm); ok {
			form := fm.getForm()
			if form != nil && form.State == huh.StateCompleted {
				//TODO form is filled out and we should take action depending on usecase
				cmds = append(cmds, fm.getResultCmd())
			}
		}
	}
	return p, tea.Batch(cmds...)
}

var (
	indigo           = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	popupBorderStyle = focusedBorderStyle.BorderForeground(indigo)
	// popupTitleStyle   = lipgloss.NewStyle().Bold(true).Foreground(indigo)
	popupMessageStyle = lipgloss.NewStyle().Italic(true).Foreground(indigo)
	// popupHelpStyle    = lipgloss.NewStyle().Foreground(gray).Italic(true)
)

// View implements tea.Model.
func (p *popup) View() string {
	// log(fmt.Sprintf("[%s] - %s", time.Now().Format(time.RFC3339Nano), "called popup.view"))
	help := "[esc][enter]"
	if len(p.message) > len(help) {
		padding := strings.Repeat(" ", len(p.message)-len(help))
		help = padding + help
	}
	popupTitleStyle := p.theme.Focused.Title
	// popupMessageStyle := p.theme.Focused.Description
	popupHelpStyle := p.theme.Help.FullDesc

	if p.nested != nil {
		return popupBorderStyle.Render(p.nested.View())
	}
	s := fmt.Sprintf("%s\n\n%s\n\n%s",
		popupTitleStyle.Render(p.title),
		popupMessageStyle.Render(p.message),
		popupHelpStyle.Render(help))

	return popupBorderStyle.Render(s)
}
