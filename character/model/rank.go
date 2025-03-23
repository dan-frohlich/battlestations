package model

type Rank int

var rankTitle = []string{
	"civilian",
	"Ensign",
	"Lt. Jr. Grade",
	"Lieutenant",
	"Commander",
	"Captain",
	"Major",
	"Colonel",
	"Commodore",
	"Admiral",
	"Fleet Admiral",
	"Commander in Chief",
	"Senator of the Republic",
}

var rankTitleAbv = []string{
	"",
	"ens",
	"lt jg",
	"lt",
	"cmdr",
	"cptn",
	"mj",
	"col",
	"com",
	"adm",
	"fl adm",
	"c n c",
	"sen",
}

func (r Rank) Title() string {
	if r < 0 || int(r) >= len(rankTitle) {
		return ""
	}
	return rankTitle[int(r)]
}

func (r Rank) TitleAbv() string {
	if r < 0 || int(r) >= len(rankTitleAbv) {
		return ""
	}
	return rankTitleAbv[int(r)]
}
