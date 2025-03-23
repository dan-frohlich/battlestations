package character

import (
	"fmt"

	"github.com/dan-frohlich/battlestations/character/model"
	"github.com/dan-frohlich/battlestations/character/print"
)

type Manager struct {
	character model.Character
}

func (m *Manager) Print() {
	pc := char2Print(m.character)
	_ = print.WritePDFFile(pc)
}

func char2Print(c model.Character) (pc print.BSChar) {
	pc.Ability = c.Species.Ability.String()
	pc.Athletics = c.Athletics.String()
	pc.BaseHP = int2Str(c.Species.BaseHT)
	pc.Carry = int2Str(c.Carry)
	pc.Combat = c.Combat.String()
	pc.Credits = int2Str(c.Credits)
	pc.Engineering = c.Engineering.String()
	pc.Equipment = make([]print.Equipment, 0, len(c.Gear))
	for _, g := range c.Gear {
		pc.Equipment = append(pc.Equipment,
			print.Equipment{Name: g.Name, Notes: g.Notes, Mass: int2Str(g.Mass), Status: "?"}) //TODO
	}
	pc.Experience = int2Str(c.Experience)
	pc.HP = int2Str(c.HP)
	pc.Hands = int2Str(c.Hands)
	pc.Luck = int2Str(c.Luck)
	pc.Move = int2Str(c.Move)
	pc.Name = c.Name
	pc.Pilot = c.Pilot.String()
	pc.Prestige = int2Str(c.Prestige)
	pc.Profession = c.Profession
	pc.Rank = int2Str(c.Rank)
	pc.Science = c.Science.String()
	pc.SpecialAbilities = make([]print.SpecialAbility, 0, len(c.SpecialAbilities))
	for _, sa := range c.SpecialAbilities {
		pool := ""
		if sa.PoolFunc != nil {
			pool = int2Str(sa.PoolFunc(c))
		}
		notes := sa.Summary
		if sa.OutputSummary != "" {
			notes = sa.OutputSummary
		}
		pc.SpecialAbilities = append(pc.SpecialAbilities, print.SpecialAbility{Name: sa.Name, Notes: notes, Pool: pool})
	}
	pc.Species = c.Species.Name
	pc.Target = int2Str(c.Species.TN)

	return pc
}

func int2Str(x int) string {
	return fmt.Sprintf("%d", x)
}
