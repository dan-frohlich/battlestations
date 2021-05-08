package main

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type bsChar map[string]interface{}

type bsChar2 struct {
	Name             string           `yaml:"name"`
	Profession       string           `yaml:"profession"`
	Athletics        string           `yaml:"athletics"`
	Combat           string           `yaml:"combat"`
	Engineering      string           `yaml:"engineering"`
	Pilot            string           `yaml:"pilot"`
	Science          string           `yaml:"science"`
	HP               string           `yaml:"hp"`
	Carry            string           `yaml:"carry"`
	Rank             string           `yaml:"rank"`
	Prestige         string           `yaml:"prestige"`
	Experience       string           `yaml:"experience"`
	Credits          string           `yaml:"credits"`
	Species          string           `yaml:"species"`
	Ability          string           `yaml:"alien_ability"`
	BaseHP           string           `yaml:"base_hp"`
	Move             string           `yaml:"move"`
	Target           string           `yaml:"target"`
	Hands            string           `yaml:"hands"`
	Luck             string           `yaml:"luck"`
	SpecialAbilities []SpecialAbility `yaml:"special_abilities"`
	Equipment        []Equipment      `yaml:"equipment"`
}

type SpecialAbility struct {
	Name  string `yaml:"name"`
	Notes string `yaml:"notes"`
	Pool  string `yaml:"pool"`
}

type Equipment struct {
	Name   string `yaml:"name"`
	Notes  string `yaml:"notes"`
	Mass   string `yaml:"mass"`
	Status string `yaml:"status"`
}

func (ch bsChar2) toChar() bsChar {
	x, _ := yaml.Marshal(ch)
	t := bsChar{}
	yaml.Unmarshal(x, &t)
	return t
}

func loadCharFromReader(r io.Reader) bsChar2 {
	t := bsChar2{}
	data, e := ioutil.ReadAll(r)
	check(e, "read char")
	e = yaml.Unmarshal(data, &t)
	check(e, "unmarshal  char")
	return t
}

func (charData bsChar2) isLarge() bool {
	if len(charData.SpecialAbilities) > 10 {
		return true
	}

	if len(charData.Equipment) > 10 {
		return true
	}
	return false
}
