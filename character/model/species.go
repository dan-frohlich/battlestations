package model

import (
	"fmt"
	"sort"
	"strings"
)

type Hands int

func (h Hands) String() string {
	if h < 0 {
		return "∞"
	}
	return fmt.Sprintf("%d", h)
}

type Species struct {
	Name    string           `yaml:"name"`
	BaseHT  int              `yaml:"base_health"`
	TN      int              `yaml:"target"`
	Hands   Hands            `yaml:"hands"`
	Move    int              `yaml:"move"`
	Armor   bool             `yaml:"armor"`
	Ability SpeciesAbilities `yaml:"species_abilities"`
}

type SpeciesAbilities []SpeciesAbility

func (sa SpeciesAbilities) String() (s string) {
	for _, ability := range sa {
		s += ability.Details()
	}
	return s
}

type SpeciesAbility struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func GetSpecies(name string) (Species, bool) {
	for _, s := range aliens {
		if strings.EqualFold(s.Name, name) {
			return s, true
		}
	}
	return Species{}, false
}

func (sa SpeciesAbility) Details() string {
	return fmt.Sprintf("%s - %s", sa.Name, sa.Description)
}

func SpeciesNames() (names []string) {
	for _, s := range aliens {
		names = append(names, s.Name)
	}
	sort.Strings(names)
	return
}

var aliens = []Species{
	{Name: "Bot (Sentient)", BaseHT: 4, TN: 8, Hands: 1, Move: 4, Armor: false,
		Ability: []SpeciesAbility{{Name: "Mechanical",
			Description: "-1 point per die of damage. 1 free upgrade per Rank (no other upgrades allowed). No cyberware. No drugs or toxins."}}},
	{Name: "Canosian", BaseHT: 8, TN: 8, Hands: -1, Move: 5, Armor: false,
		Ability: []SpeciesAbility{{Name: "Tumble:",
			Description: "You get a second move action either before or after your action."}}},
	{Name: "Human", BaseHT: 5, TN: 8, Hands: 2, Move: 5, Armor: true,
		Ability: []SpeciesAbility{{Name: "Willpower:",
			Description: "You may reroll one or both dice when making a skill check within your profession. You need not decide whether to reroll both before rerolling either."}}},
	{Name: "Silicoid", BaseHT: 9, TN: 7, Hands: 1, Move: 4, Armor: false,
		Ability: []SpeciesAbility{
			{Name: "Rocky", Description: "Each time you would be dealt damage reduce it by 1d6."},
			{Name: "Strong", Description: "Add +10 Carry and +1 point of damage dealt with melee attacks."}}},
	{Name: "Tentac", BaseHT: 6, TN: 9, Hands: -1, Move: 6, Armor: false,
		Ability: []SpeciesAbility{{Name: "Resilient",
			Description: "Roll an extra die each time you are damaged and remove the highest die."}}},
	{Name: "Xeloxian", BaseHT: 5, TN: 8, Hands: 6, Move: 2, Armor: true,
		Ability: []SpeciesAbility{
			{Name: "Fistwalk", Description: "Add free hands to move value."},
			{Name: "Aggressive", Description: "Add +1 point to direct personal attack damage."}}},
	{Name: "Zoallan", BaseHT: 4, TN: 9, Hands: 3, Move: 7, Armor: false,
		Ability: []SpeciesAbility{{Name: "Carapace",
			Description: "-2 points of damage from all sources."}}},
	{Name: "Blootian", BaseHT: 0, TN: 7, Hands: 5, Move: 6, Armor: false,
		Ability: []SpeciesAbility{{Name: "Bubbly",
			Description: "Free Reroll on Life Support Checks and the largest die in an attack (or a needler hit) pops a limb bubble instead of dealing damage. Add limb bubbles to move and hands attribute. Subtract limb bubbles from Target #. Athletics check of 8 to regrow limb bubble or when healing regrow one limb bubble in place of a die of hit point healing. Begin play with maximum bubbles (4)."}}},
	{Name: "Diploid", BaseHT: 4, TN: 8, Hands: 2, Move: 3, Armor: true,
		Ability: []SpeciesAbility{{Name: "Bifurcation",
			Description: "You get two separate phases each phase as if you were two characters with the same body. All active skill check actions in phases (not upgrades, requisitions) are at a +3 difficulty penalty."}}},
	{Name: "Fungaloid", BaseHT: 8, TN: 9, Hands: 2, Move: 4, Armor: false,
		Ability: []SpeciesAbility{{Name: "Regenerate",
			Description: "Recover 2 points at the end of each phase as long as damage isn't more than your Hit Points."}}},
	{Name: "Kerbite", BaseHT: 5, TN: 9, Hands: 5, Move: 6, Armor: true,
		Ability: []SpeciesAbility{{Name: "Cooperative",
			Description: "Once per phase, you get a free Assist action on a friendly as you are moving adjacent to them during your move action."}}},
	{Name: "Trundlian", BaseHT: 0, TN: 8, Hands: 7, Move: 2, Armor: false,
		Ability: []SpeciesAbility{{Name: "Versatile",
			Description: "You may reroll “1”'s in your initial skill checks. Add your empty hands to your Hit Points or move in any combination at any time."}}},
	{Name: "Vomeg", BaseHT: 7, TN: 7, Hands: 3, Move: 5, Armor: true,
		Ability: []SpeciesAbility{{Name: "Reach",
			Description: "You may act from any adjacent square as if you were in it. Doing so does not provoke free attacks and is at no penalty for being occupied."}}},
	{Name: "Whistler", BaseHT: 9, TN: 7, Hands: 4, Move: 5, Armor: false,
		Ability: []SpeciesAbility{{Name: "Puff",
			Description: "One of your moves in each move action may be a jet move that doesn't require a skill check."}}},
}
