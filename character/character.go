package character

import (
	"io"

	"gopkg.in/yaml.v2"
)

type Character struct {
	Name             string           `yaml:"name"`
	Profession       Profession       `yaml:"profession"`
	Athletics        SkillLevel       `yaml:"athletics"`
	Combat           SkillLevel       `yaml:"combat"`
	Engineering      SkillLevel       `yaml:"engineering"`
	Pilot            SkillLevel       `yaml:"pilot"`
	Science          SkillLevel       `yaml:"science"`
	Diplomacy        SkillLevel       `yaml:"diplomacy"`
	Prionic          SkillLevel       `yaml:"psionic"`
	HP               int              `yaml:"hp"`
	Carry            Mass             `yaml:"carry"`
	Rank             Rank             `yaml:"rank"`
	Prestige         Prestige         `yaml:"prestige"`
	Experience       Experience       `yaml:"experience"`
	Credits          Credits          `yaml:"credits"`
	Species          Species          `yaml:"species"`
	Luck             Luck             `yaml:"luck"`
	SpecialAbilities []SpecialAbility `yaml:"special_abilities"`
	Equipment        []Equipment      `yaml:"equipment"`
}

type Credits int
type Experience int
type Luck int
type Mass int
type Prestige int
type SkillLevel int

type charMap map[string]string

func (ch Character) toMap() charMap {
	x, _ := yaml.Marshal(ch)
	t := charMap{}
	yaml.Unmarshal(x, &t)
	return t
}

func LoadCharFromReader(r io.Reader) (Character, error) {
	t := Character{}
	data, e := io.ReadAll(r)
	if e != nil {
		return t, e
	}
	e = yaml.Unmarshal(data, &t)
	if e != nil {
		return t, e
	}
	return t, nil
}

func (ch Character) isLarge() bool {
	if len(ch.SpecialAbilities) > 10 {
		return true
	}

	if len(ch.Equipment) > 10 {
		return true
	}
	return false
}
