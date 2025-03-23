package model

type Character struct {
	Name             string           `yaml:"name"`
	Player           string           `yaml:"player"`
	Profession       string           `yaml:"profession"`
	Athletics        int              `yaml:"athletics"`
	Combat           int              `yaml:"combat"`
	Engineering      int              `yaml:"engineering"`
	Pilot            int              `yaml:"pilot"`
	Science          int              `yaml:"science"`
	Diplomacy        int              `yaml:"diplomacy"`
	Psionics         int              `yaml:"psionics"`
	Sanity           int              `yaml:"sanity"`
	HP               int              `yaml:"hp"`
	Carry            int              `yaml:"carry"`
	Rank             int              `yaml:"rank"`
	Prestige         int              `yaml:"prestige"`
	Experience       int              `yaml:"experience"`
	Credits          int              `yaml:"credits"`
	Move             int              `yaml:"move"`
	Target           int              `yaml:"target"`
	Hands            int              `yaml:"hands"`
	Luck             int              `yaml:"luck"`
	Species          Species          `yaml:"species"`
	SpecialAbilities []SpecialAbility `yaml:"special_abilities"`
	Gear             []Gear           `yaml:"Gear"`
}
