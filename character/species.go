package character

type Species struct {
	Name      string           `yaml:"name"`
	BaseHP    int              `yaml:"base_hp"`
	Move      int              `yaml:"move"`
	TN        int              `yaml:"target"`
	Hands     int              `yaml:"hands"`
	Armor     YN               `yaml:"armor"`
	Abilities []SpeciesAbility `yaml:"abilities"`
}

type SpeciesAbility struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type YN bool

const (
	Y YN = true
	N YN = false
)

func (v YN) String() string {
	if v {
		return "y"
	}
	return "n"
}

var (
	CanosianSpecies Species = Species{
		Name: "Canosian", BaseHP: 8, TN: 8, Hands: -1, Move: 5, Armor: N,
		Abilities: []SpeciesAbility{{Name: "Tumble:",
			Description: "You get a second move action either before or after your action."}}}
)
