package character

import (
	"fmt"
	"strings"

	"github.com/dan-frohlich/battlestations/character/model"
	"github.com/dan-frohlich/battlestations/character/print"
)

type Manager struct {
	character model.Character
}

func (m *Manager) SetCharacter(c model.Character) {
	m.character = c
}

func (m *Manager) Print() {
	pc := char2Print(m.character)
	_ = print.WritePDFFile(pc)
}

func char2Print(c model.Character) (pc print.BSChar) {
	pc.Ability = c.Species.Abilities.String()
	pc.Athletics = c.Athletics.String()
	pc.BaseHP = int2Str(c.Species.BaseHT)
	pc.Carry = int2Str(c.Carry())
	pc.Combat = c.Combat.String()
	pc.Credits = int2Str(c.Credits)
	pc.Engineering = c.Engineering.String()
	pc.Equipment = make([]print.Equipment, 0, len(c.Gear))
	for _, g := range c.Gear {
		var status []string
		if g.Installed {
			status = append(status, "instld")
		} else if g.Equiped {
			status = append(status, "equip")
		} else {
			status = append(status, "stwd")
		}
		if g.Upgraded {
			status = append(status, "upgrd")
		}
		notes := g.Notes
		if g.OutputNotes != "" {
			notes = g.OutputNotes
		}
		pc.Equipment = append(pc.Equipment,
			print.Equipment{Name: g.Name, Notes: notes, Mass: int2Str(g.Mass), Status: strings.Join(status, " / ")})
	}
	pc.Experience = int2Str(c.Experience)
	pc.HP = int2Str(c.HP())
	pc.Hands = int2Str(int(c.Species.Hands))
	pc.Luck = int2Str(c.Luck())
	pc.Move = int2Str(c.Species.Move + c.SpecialAbilities.MoveBonus())
	pc.Name = c.Name
	pc.Pilot = c.Pilot.String()
	pc.Prestige = int2Str(c.Prestige)
	pc.Profession = c.Profession

	// if we have one or more optionl skills, display best one
	names := map[string]int{"Diplomacy": int(c.Diplomacy), "Psionics": int(c.Psionics), "Sanity": int(c.Sanity)}
	maxVal := 0
	maxName := ""
	OptionalSkills := 0
	for k, v := range names {
		if v > 0 {
			OptionalSkills++
		}
		if v > maxVal {
			maxVal = v
			maxName = k
		}
	}
	if maxVal > 0 {
		pc.TinyNote = fmt.Sprintf("%s: %d", maxName, maxVal)
	}
	// if OptionalSkills > 1 {
	//uh-oh, not sure what to do with the other ones... TODO: FIXME
	// }

	pc.Rank = int2Str(int(c.Rank)) + " " + c.Rank.TitleAbv()
	pc.Science = c.Science.String()
	pc.SpecialAbilities = make([]print.SpecialAbility, 0, len(c.SpecialAbilities))
	for _, sa := range c.SpecialAbilities {
		pool := ""
		// if sa.PoolFunc != nil {
		// 	pool = int2Str(sa.PoolFunc(c))
		// }
		if sa.Pool != "" {
			i, _ := sa.Pool.Calculate(c)
			pool = int2Str(i)
		} else {
			pool = "-"
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
