package character

type SpecialAbility struct {
	Name  string
	Notes string
	Pool  PoolCalculator
}

type PoolCalculator func(c Character) int

var BosunSpecialAbility = SpecialAbility{
	Name:  "Bosun",
	Notes: "+1 Remote. BS reroll pool",
	Pool:  func(c Character) int { return 5 },
}
