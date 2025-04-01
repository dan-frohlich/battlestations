package tui

type panel int

const (
	_ panel = iota
	menuPanel
	detailPanel
	statusPanel
)

func (p panel) String() string {
	switch p {
	case menuPanel:
		return "menu_panel"
	case detailPanel:
		return "detail_panel"
	case statusPanel:
		return "status_panel"
	}
	return "unknown_panel"
}

func (p panel) next() panel {
	switch p {
	case statusPanel:
		return menuPanel
	default:
		return p + 1
	}
}

func (p panel) prev() panel {
	switch p {
	case menuPanel:
		return statusPanel
	default:
		return p - 1
	}
}
