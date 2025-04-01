package tui

type usecaseView int

const (
	_ usecaseView = iota
	mainView
	loadCharSubView
	quitSubView
	newCharView
	setStatsSubView
	setSpeciesSubView
	setNameSubView
	setProfessionSubView
	setBasicGearSubView
	manageCharView
	saveCharSubView
	charPreviewSubView
	printCharSubView
	missionAftermathView
	purchaseGearSubView
	addPrestigeSubView
	addExpSubView
	addCreditsSubView
	requisitionSubView
	pillageSubView
)

func (u usecaseView) String() string {
	switch u {
	case mainView:
		return "main_usecase"
	case loadCharSubView:
		return "load_char_usecase"
	case newCharView:
		return "new_char_usecase"
	case setStatsSubView:
		return "set_stats_usecase"
	case setSpeciesSubView:
		return "set_species_usecase"
	case setNameSubView:
		return "set_name_usecase"
	case setProfessionSubView:
		return "set_profession_usecase"
	case setBasicGearSubView:
		return "set_basic_gear_usecase"
	case manageCharView:
		return "manage_char_usecase"
	case saveCharSubView:
		return "save_char_usecase"
	case charPreviewSubView:
		return "char_preview_usecase"
	case printCharSubView:
		return "print_char_usecase"
	case missionAftermathView:
		return "mission_aftermath_usecase"
	case addPrestigeSubView:
		return "add_prestige_usecase"
	case addExpSubView:
		return "add_exp_usecase"
	case addCreditsSubView:
		return "add_credits_usecase"
	case requisitionSubView:
		return "requisition_usecase"
	case pillageSubView:
		return "pillage_usecase"
	default:
		return "unknown_usecase"
	}
}

// valid transitions
var _z = struct{}{}
var _ = map[usecaseView]map[usecaseView]struct{}{
	mainView: {
		loadCharSubView: _z,
		newCharView:     _z,
		quitSubView:     _z,
	},
	loadCharSubView: {
		newCharView:    _z,
		manageCharView: _z,
	},
	newCharView: {
		mainView:             _z,
		setStatsSubView:      _z,
		setSpeciesSubView:    _z,
		setNameSubView:       _z,
		setProfessionSubView: _z,
		setBasicGearSubView:  _z,
	},
	setStatsSubView: {
		newCharView:    _z,
		manageCharView: _z,
	},
	setSpeciesSubView: {
		newCharView:    _z,
		manageCharView: _z,
	},
	setNameSubView: {
		newCharView:    _z,
		manageCharView: _z,
	},
	setProfessionSubView: {
		newCharView:    _z,
		manageCharView: _z,
	},
	setBasicGearSubView: {
		newCharView:    _z,
		manageCharView: _z,
	},
	manageCharView: {
		mainView:             _z,
		charPreviewSubView:   _z,
		printCharSubView:     _z,
		saveCharSubView:      _z,
		missionAftermathView: _z,
		purchaseGearSubView:  _z,
	},
	charPreviewSubView: {
		manageCharView: _z,
	},
	printCharSubView: {
		manageCharView: _z,
	},
	saveCharSubView: {
		manageCharView: _z,
	},
	purchaseGearSubView: {
		manageCharView: _z,
	},
	missionAftermathView: {
		manageCharView:     _z,
		addPrestigeSubView: _z,
		addExpSubView:      _z,
		addCreditsSubView:  _z,
		requisitionSubView: _z,
		pillageSubView:     _z,
	},
}
