package manager

import (
	"fmt"

	"github.com/dan-frohlich/battlestations/character"
	"github.com/dan-frohlich/battlestations/print"
)

func convertForPrinting(c character.Character) print.BSChar {

	sa := ""
	cr := ""
	for i, ssa := range c.Species.Abilities {
		if i > 0 {
			cr = "\n"
		}
		sa = fmt.Sprintf("%s%s%s: %s", sa, cr, ssa.Name, ssa.Description)
	}
	return print.BSChar{
		Name:        c.Name,
		Profession:  c.Profession.String(),
		Athletics:   c.Athletics.String(),
		Combat:      c.Combat.String(),
		Engineering: c.Engineering.String(),
		Pilot:       c.Pilot.String(),
		Science:     c.Science.String(),
		HP:          fmt.Sprintf("%d", c.Rank.Int()+c.Species.BaseHP),
		Carry:       fmt.Sprintf("%d", c.Athletics.Int()*10),
		Rank:        fmt.Sprintf("%d / %s", c.Rank.Int(), c.Rank.AbrvString()),
		Prestige:    c.Prestige.String(),
		Experience:  c.Experience.String(),
		Credits:     c.Credits.String(),
		Species:     c.Species.Name,
		Ability:     sa,
		BaseHP:      fmt.Sprintf("%d", c.Species.BaseHP),
		Move:        fmt.Sprintf("%d", c.Species.Move),
		Target:      fmt.Sprintf("%d", c.Species.TN),
		Hands:       c.Species.Hands.String(),
		Luck:        fmt.Sprintf("%d", c.Rank.Int()+5),
	}
}
