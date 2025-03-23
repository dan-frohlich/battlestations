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
	Profession       string             `yaml:"profession"`
	Athletics        SkillLevel         `yaml:"athletics"`
	Combat           SkillLevel         `yaml:"combat"`
	Engineering      SkillLevel         `yaml:"engineering"`
	Pilot            SkillLevel         `yaml:"pilot"`
	Science          SkillLevel         `yaml:"science"`
	Diplomacy        OptionalSkillLevel `yaml:"diplomacy,omitempty"`
	Psionics         OptionalSkillLevel `yaml:"psionics,omitempty"`
	Sanity           OptionalSkillLevel `yaml:"sanity,omitempty"`
	HP               int                `yaml:"hp"`
	Carry            int                `yaml:"carry"`
	Rank             int                `yaml:"rank"`
	Prestige         int                `yaml:"prestige"`
	Experience       int                `yaml:"experience"`
	Credits          int                `yaml:"credits"`
	Move             int                `yaml:"move"`
	Target           int                `yaml:"target"`
	Hands            int                `yaml:"hands"`
	Luck             int                `yaml:"luck"`
	Species          Species            `yaml:"species"`
	SpecialAbilities []SpecialAbility   `yaml:"special_abilities"`
	Gear             []Gear             `yaml:"gear"`
	GearRef          []GearRef          `yaml:"gear_status"`
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

func hydrate(c Character) Character {
	for i := range c.SpecialAbilities {
		sa := GetAbility(c.SpecialAbilities[i].Name)
		c.SpecialAbilities[i].Summary = sa.Summary
		c.SpecialAbilities[i].FullDescription = sa.FullDescription
		if c.SpecialAbilities[i].OutputSummary != "" { // if not overridden in char save
			c.SpecialAbilities[i].OutputSummary = sa.OutputSummary
		}
		c.SpecialAbilities[i].Types = sa.Types
		c.SpecialAbilities[i].PoolFunc = sa.PoolFunc
	}
	return c
}
