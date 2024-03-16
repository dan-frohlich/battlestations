package character

import "fmt"

type Species struct {
	Name      string           `yaml:"name"`
	BaseHP    int              `yaml:"base_hp"`
	Move      int              `yaml:"move"`
	TN        int              `yaml:"target"`
	Hands     Hands            `yaml:"hands"`
	Armor     YN               `yaml:"armor"`
	Abilities []SpeciesAbility `yaml:"abilities"`
}

type SpeciesAbility struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Hands int

func (h Hands) Int() int { return int(h) }

func (h Hands) String() string {
	if h < 0 {
		return "OO"
	}
	return fmt.Sprintf("%d", h)
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
		Abilities: []SpeciesAbility{{Name: "Tumble",
			Description: "You get a second move action either before or after your action."}}}
	XeloxianSpecies Species = Species{
		Name: "Xeloxian", BaseHP: 5, TN: 8, Hands: 6, Move: 2, Armor: Y,
		Abilities: []SpeciesAbility{
			{Name: "Fistwalk", Description: "Add free hands to move value."},
			{Name: "Aggressive", Description: "Add +1 point to direct personal attack damage."}}}
)

/*
	{Name: "Xeloxian", BaseHT: 5, TN: 8, Hands: 6, Move: 2, Armor: Y,
		Ability: []SpeciesAbility{
			{Name: "Fistwalk", Description: "Add free hands to move value."},
			{Name: "Aggressive", Description: "Add +1 point to direct personal attack damage."}}},

*/
