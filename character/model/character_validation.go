package model

import (
	"fmt"
	"sort"
	"strings"
)

type ValidationMessage string

// Error error interface
func (vm ValidationMessage) Error() string {
	return string(vm)
}

// String stringer interface
func (vm ValidationMessage) String() string {
	return string(vm)
}

// ID return val;idation id
func (vm ValidationMessage) ID() string {
	if len(vm) > 0 {
		return strings.Split(string(vm), " ")[0]
	}
	return "?"
}

func ValidationMessagef(format string, args ...any) ValidationMessage {
	return ValidationMessage(fmt.Sprintf(format, args...))
}

type CharacterValidatorFunc func(id string, c Character) (issues []ValidationMessage)

type CharacterValidator struct {
	validators map[string]CharacterValidatorFunc
}

func (cv *CharacterValidator) Register(id string, f CharacterValidatorFunc) (err error) {
	if _, ok := cv.validators[id]; ok {
		return fmt.Errorf("")
	}
	return nil
}

func (cv *CharacterValidator) ValidateAll(c Character) (issues []ValidationMessage, errs []error) {
	var ids []string
	for id := range cv.validators {
		ids = append(ids, id)
	}

	sort.Strings(ids)
	for _, id := range ids {
		if f := cv.validators[id]; f != nil {
			issues = append(issues, f(id, c)...)
		} else {
			errs = append(errs, fmt.Errorf("not func registered to [%s]", id))
		}
	}
	return issues, errs
}

func NewCharacterValidator() *CharacterValidator {
	var m = map[string]CharacterValidatorFunc{
		"v01": func(id string, c Character) (issues []ValidationMessage) {
			//must have starting skill
			if len(c.StartingSkillSet.AsLevels()) < 4 {
				return []ValidationMessage{ValidationMessagef("%s - invalid starting skill set: [%s]", id, c.StartingSkillSet)}
			}
			return nil
		},
		"v02": func(id string, c Character) (issues []ValidationMessage) {
			//must not have negative attributes
			m2 := map[string]int{
				"athletics":   c.Athletics.AsInt(),
				"diplomacy":   c.Diplomacy.AsInt(),
				"engineering": c.Engineering.AsInt(),
				"pilot":       c.Pilot.AsInt(),
				"psioonics":   c.Psionics.AsInt(),
				"science":     c.Science.AsInt(),
				"sanity":      c.Sanity.AsInt(),
				"rank":        c.Rank.AsInt(),
				"prestige":    c.Prestige,
				"experience":  c.Experience,
				"credits":     c.Credits,
			}
			var keys []string
			for k := range m2 {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				if m2[k] < 0 {
					issues = append(issues, ValidationMessagef("%s - trait %s can not be negative: %d", id, k, m2[k]))
				}
			}
			return issues
		},
		"v03": func(id string, c Character) (issues []ValidationMessage) {
			//must have special abilities equal to rank
			if len(c.SpecialAbilities) != c.Rank.AsInt() {
				return []ValidationMessage{ValidationMessagef("%s - should have 1 ability per rank: expected %d abilities but found %d", id, c.Rank, len(c.SpecialAbilities))}
			}
			return nil
		},
		"v04": func(id string, c Character) (issues []ValidationMessage) {
			//must not have unspent prestige
			if c.Prestige >= c.Rank.AsInt()*100 {
				return []ValidationMessage{ValidationMessagef("%s - must not have unspent prestige: expected no more than %d prestige but found %d", id, c.Rank*100, c.Prestige)}
			}
			return nil
		},
		"v05": func(id string, c Character) (issues []ValidationMessage) {
			//must not have unspent prestige
			if c.Overburdened() {
				return []ValidationMessage{ValidationMessagef("%s - Overburdended: carrying %d with a carry limit of %d", id, c.Load(), c.Carry())}
			}
			return nil
		},
	}
	cv := &CharacterValidator{
		validators: m,
	}
	return cv
}
