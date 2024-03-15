package character

import "gopkg.in/yaml.v2"

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
