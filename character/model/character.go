package model

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type SkillLevel int

func (sl SkillLevel) String() string {
	return fmt.Sprintf("%d", sl)
}

type OptionalSkillLevel int

func (sl OptionalSkillLevel) String() string {
	if sl == 0 {
		return ""
	}
	return fmt.Sprintf("%d", sl)
}

type Character struct {
	Name             string             `yaml:"name"`
	Player           string             `yaml:"player"`
	StartingSkillSet string             `yaml:"starting_skill_set"`
	Profession       string             `yaml:"profession"`
	Athletics        SkillLevel         `yaml:"athletics"`
	Combat           SkillLevel         `yaml:"combat"`
	Engineering      SkillLevel         `yaml:"engineering"`
	Pilot            SkillLevel         `yaml:"pilot"`
	Science          SkillLevel         `yaml:"science"`
	Diplomacy        OptionalSkillLevel `yaml:"diplomacy,omitempty"`
	Psionics         OptionalSkillLevel `yaml:"psionics,omitempty"`
	Sanity           OptionalSkillLevel `yaml:"sanity,omitempty"`
	Rank             Rank               `yaml:"rank"`
	Prestige         int                `yaml:"prestige"`
	Experience       int                `yaml:"experience"`
	Credits          int                `yaml:"credits"`
	Move             int                `yaml:"move"`
	Target           int                `yaml:"target"`
	Hands            int                `yaml:"hands"`
	Species          Species            `yaml:"species"`
	SpecialAbilities []SpecialAbility   `yaml:"special_abilities"`
	Gear             []GearRef          `yaml:"gear"`
}

func (c Character) HP() int {
	return int(c.Athletics) + int(c.Rank) + c.Species.BaseHT
}

func (c Character) Luck() int {
	return int(c.Rank) + 5
}

func (c Character) Carry() int {
	return int(c.Athletics) * 10
}

func MustLoadCharacter(data []byte) Character {
	c, err := LoadCharacter(data)
	if err != nil {
		panic(err)
	}
	return c
}

func LoadCharacter(data []byte) (c Character, err error) {
	err = yaml.Unmarshal(data, &c)
	c = hydrate(c)

	return c, err
}

// ItemNotFoundString an item name ending with this string was not found in the BattleStations DB
const ItemNotFoundString = "*"

func hydrate(c Character) Character {
	c.Name += ItemNotFoundString //TODO remove after debugging

	if s, ok := GetSpecies(c.Species.Name); ok {
		c.Species.Armor = s.Armor
		c.Species.BaseHT = s.BaseHT
		c.Species.TN = s.TN
		c.Species.Hands = s.Hands
		c.Species.Move = s.Move
		//index species abilities from DB
		m := make(map[string]SpeciesAbility)
		for _, sa := range s.Abilities {
			m[sa.Name] = sa
		}
		if len(c.Species.Abilities) == 0 {
			//if we dont have any, copy from Battlestations DB
			c.Species.Abilities = s.Abilities
		} else {
			//if we do have abilities, hydrate them from Battlestations DB
			for i := range c.Species.Abilities {
				if sa, ok := m[c.Species.Abilities[i].Name]; ok {
					c.Species.Abilities[i].Description = sa.Description
					if c.Species.Abilities[i].OutputDescription == "" { // if not overridden in char save
						c.Species.Abilities[i].OutputDescription = sa.OutputDescription
					}
				}
			}
		}
	} else {
		c.Species.Name += ItemNotFoundString
	}
	for i := range c.SpecialAbilities {
		sa := GetAbility(c.SpecialAbilities[i].Name)
		if sa.Name == "" {
			c.SpecialAbilities[i].Name += ItemNotFoundString
			continue
		}
		c.SpecialAbilities[i].Summary = sa.Summary
		c.SpecialAbilities[i].FullDescription = sa.FullDescription
		if c.SpecialAbilities[i].OutputSummary == "" { // if not overridden in char save
			c.SpecialAbilities[i].OutputSummary = sa.OutputSummary
		}
		c.SpecialAbilities[i].Types = sa.Types
		c.SpecialAbilities[i].PoolCode = sa.PoolCode
	}

	for i := range c.Gear {
		g := GetGear(c.Gear[i].Name)
		if g.Name == "" {
			c.Gear[i].Name += ItemNotFoundString
			continue
		}
		c.Gear[i].Notes = g.Notes
		c.Gear[i].Cost = g.Cost
		c.Gear[i].Energy = g.Energy
		if !c.Gear[i].Upgraded && c.Gear[i].Mass == 0 {
			c.Gear[i].Mass = g.Mass
		}
		c.Gear[i].Notes = g.Notes
		if c.Gear[i].OutputNotes == "" { // if not overridden in char save
			c.Gear[i].OutputNotes = g.OutputNotes
		}
		c.Gear[i].Type = g.Type
	}
	return c
}
