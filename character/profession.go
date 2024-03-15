package character

type Profession string

func (p Profession) String() string { return string(p) }

const (
	EngineerProfession  Profession = "Engineer"
	MarineProfession    Profession = "Marine"
	PilotProfession     Profession = "Pilot"
	ScientistProfession Profession = "Scientist"
	AthleetProfession   Profession = "Athleet"
	DiplomatProfession  Profession = "Diplomat"
	PsionProfession     Profession = "Psion"
)

var (
	AltProfessions  = []Profession{AthleetProfession, DiplomatProfession, PsionProfession}
	CoreProfessions = []Profession{EngineerProfession, MarineProfession, PilotProfession, ScientistProfession}
	AllProfessions  = []Profession{AthleetProfession, DiplomatProfession, PsionProfession, EngineerProfession, MarineProfession, PilotProfession, ScientistProfession}
)
