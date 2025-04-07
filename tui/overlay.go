package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

var _ tea.Model = &olay{}

type olay struct {
	displayOverlay bool
	winHeight      int
	fg             *popup
	bg             tea.Model
	overlay        *overlay.Model
	theme          *huh.Theme
}

func newOverlay(bg tea.Model, theme *huh.Theme) *olay {
	note := huh.NewNote().Title("Note Title").Description("a messge to display")
	fg := &popup{
		title:   "Popup Title",
		message: "popup message to display",
		note:    note,
		theme:   theme,
		form:    huh.NewForm(huh.NewGroup(note).WithTheme(theme)).WithTheme(theme).WithShowHelp(true),
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
	// log(fmt.Sprintf("[%s] - called olay.update([%[2]T][%[2]s])", time.Now().Format(time.RFC3339Nano), msg))
	var cmds []tea.Cmd
	//cmds = append(cmds, makeStatusCmd(debugLevel, fmt.Sprintf("overlay: [%T[1] - %[1]s", stringMsg(msg))))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// log(fmt.Sprintf("[%s] - called olay.update([%[2]T][%[2]s])", time.Now().Format(time.RFC3339Nano), msg))
		// cmds = append(cmds, makeStatusCmd(infoLevel, fmt.Sprintf("overlay received: [%T[1] - %[1]s", stringMsg(msg))))
		if o.displayOverlay {
			// what key was pressed?
			switch msg.String() {
			case "esc", "enter":
				o.displayOverlay = false
			}
			return o, nil
		}
	case tea.WindowSizeMsg:
		o.winHeight = msg.Height
	case popupMsg:
		o.displayOverlay = true
		l, cmd := o.fg.Update(msg)
		if ll, ok := l.(*popup); ok {
			o.fg = ll
		}
		return o, cmd
	}

	l, cmd := o.overlay.Update(msg)
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

type popup struct {
	title   string
	message string
	// help    string
	theme *huh.Theme
	form  *huh.Form
	note  *huh.Note
}

// Init implements tea.Model.
func (p *popup) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (p *popup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case popupMsg:
		p.title = msg.title
		p.message = msg.message
	}
	return p, nil
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

	s := fmt.Sprintf("%s\n\n%s\n\n%s",
		popupTitleStyle.Render(p.title),
		popupMessageStyle.Render(p.message),
		popupHelpStyle.Render(help))

	return popupBorderStyle.Render(s)
}
