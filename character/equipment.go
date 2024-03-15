package character

import "gopkg.in/yaml.v2"

type Equipment struct {
	Name   string `yaml:"name"`
	Notes  string `yaml:"notes"`
	Mass   int    `yaml:"mass"`
	Status string `yaml:"status"`
}

func (eq Equipment) toMap() charMap {
	b, _ := yaml.Marshal(eq)
	m := charMap{}
	yaml.Unmarshal(b, m)
	return m
}
