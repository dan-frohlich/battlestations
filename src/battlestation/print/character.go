package main

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type bsChar struct {
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

type charMap map[string]string

type SpecialAbility struct {
	Name  string `yaml:"name"`
	Notes string `yaml:"notes"`
	Pool  string `yaml:"pool"`
}

func (sa SpecialAbility) toMap() charMap {
	b, _ := yaml.Marshal(sa)
	m := charMap{}
	yaml.Unmarshal(b, m)
	return m
}

type Equipment struct {
	Name   string `yaml:"name"`
	Notes  string `yaml:"notes"`
	Mass   string `yaml:"mass"`
	Status string `yaml:"status"`
}

func (eq Equipment) toMap() charMap {
	b, _ := yaml.Marshal(eq)
	m := charMap{}
	yaml.Unmarshal(b, m)
	return m
}

func (ch bsChar) toMap() charMap {
	x, _ := yaml.Marshal(ch)
	t := charMap{}
	yaml.Unmarshal(x, &t)
	return t
}

func loadCharFromReader(r io.Reader) bsChar {
	t := bsChar{}
	data, e := ioutil.ReadAll(r)
	check(e, "read char")
	e = yaml.Unmarshal(data, &t)
	check(e, "unmarshal  char")
	return t
}

func (ch bsChar) isLarge() bool {
	if len(ch.SpecialAbilities) > 10 {
		return true
	}

	if len(ch.Equipment) > 10 {
		return true
	}
	return false
}
