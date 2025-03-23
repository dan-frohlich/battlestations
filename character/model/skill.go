package model

type SkillSet string

func (ss SkillSet) AsLevels() []SkillLevel {
	if l, ok := skillSetMap[ss]; ok {
		return l
	}
	return nil
}

func GetStatingSkillSets() []string {
	return []string{"41110", "33100", "32210", "22221", "321111"}
}

var (
	skillSetMap = map[SkillSet][]SkillLevel{
		"41110":  []SkillLevel{4, 1, 1, 1, 0},
		"33100":  []SkillLevel{3, 3, 1, 0, 0},
		"32210":  []SkillLevel{3, 2, 2, 1, 0},
		"22221":  []SkillLevel{2, 2, 2, 2, 1},
		"321111": []SkillLevel{3, 2, 1, 1, 1, 1},
	}
)
