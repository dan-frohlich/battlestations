package manager

import (
	"fmt"
	"sort"

	"github.com/dan-frohlich/battlestations/character"
)

/*
 * 1. Start at Rank 1
 * 2. Pick a species
 * 3. Select a profession
 * 4. Assign skill numbers
 * 5. Select a starting Special Ability
 * 6. Calculate Hit Points
 * 7. Get starting Equipment

 *    A starting character will be given a blaster and their choice of a MedKit or ToolKit.
 *    Characters that wear armor will also be issued armor.
 *    A one time bonus of 500 credits is given to each starting character.
 */

type SkillArray []character.SkillLevel

func (sa SkillArray) Equals(val SkillArray) bool {
	for i := range sa {
		if sa[i] != val[i] {
			return false
		}
	}
	return true
}

func (sa SkillArray) IsValid() bool {
	for i := range skillArrays {
		if sa.Equals(skillArrays[i]) {
			return true
		}
	}
	return false
}

var skillArrays = []SkillArray{
	{4, 2, 0, 0, 0},
	{4, 1, 1, 1, 0},
	{3, 3, 1, 0, 0},
	{3, 2, 2, 1, 0},
	{2, 2, 2, 2, 1},
	{3, 2, 1, 1, 1},
}

type SkillSelection struct {
	Athletics   character.SkillLevel
	Combat      character.SkillLevel
	Engineering character.SkillLevel
	Pilot       character.SkillLevel
	Science     character.SkillLevel
}

func (ss SkillSelection) String() string {
	return fmt.Sprintf("[A: %d, C: %d, E: %d, P: %d, S: %d ]", ss.Athletics, ss.Combat, ss.Engineering, ss.Pilot, ss.Science)
}

func (ss SkillSelection) toSkillArray() SkillArray {
	sa := SkillArray{ss.Athletics, ss.Combat, ss.Engineering, ss.Pilot, ss.Science}
	sort.Slice(sa, func(i, j int) bool { return int(sa[i]) > int(sa[j]) })
	return sa
}

type characterCreationStage int

const (
	_ characterCreationStage = iota
	startStage
	speciesStage
	professionStage
	assignSkillStage
	assignSpecialAbilityStage
	selectMedKitOrToolkit
	purchaseEquipmentStage
)

var stageNames = map[characterCreationStage]string{
	startStage:                "start",
	speciesStage:              "choose species",
	professionStage:           "choose profession",
	assignSkillStage:          "assign skills",
	assignSpecialAbilityStage: "choose special ability",
	selectMedKitOrToolkit:     "select kit",
	purchaseEquipmentStage:    "purchase other gear",
}

func (ccs characterCreationStage) String() string {
	if s, ok := stageNames[ccs]; ok {
		return s
	}
	return fmt.Sprintf("unknown stage: %d", ccs)
}

type CharacterCreator struct {
	stage characterCreationStage
	char  character.Character
}

func NewCharacterCreator() *CharacterCreator {
	return &CharacterCreator{
		char:  character.Character{Rank: character.EnsignRank},
		stage: speciesStage,
	}
}

func (cc *CharacterCreator) PreviousStage() {
	if cc.stage-1 == selectMedKitOrToolkit {
		//undo the gear.
		cc.char.Equipment = nil
		cc.char.Credits = 0
	}
	cc.stage--
}

func (cc *CharacterCreator) SelectSpecies(v character.Species) error {
	if cc.stage != speciesStage {
		return fmt.Errorf("can't select species in stage %s", cc.stage)
	}
	cc.char.Species = v
	cc.stage++
	return nil
}

func (cc *CharacterCreator) SelectProfession(v character.Profession) error {
	if cc.stage != professionStage {
		return fmt.Errorf("can't select profession in stage %s", cc.stage)
	}
	cc.char.Profession = v
	cc.stage++
	return nil
}

func (cc *CharacterCreator) SetName(v string) {
	cc.char.Name = v
}

func (cc *CharacterCreator) GetSkillArrays() []SkillArray {
	return skillArrays
}

func (cc *CharacterCreator) SelectSkills(v SkillSelection) error {
	if cc.stage != assignSkillStage {
		return fmt.Errorf("can't select skills in stage %s", cc.stage)
	}
	if !v.toSkillArray().IsValid() {
		return fmt.Errorf("invalid skill selection %v", v)
	}
	cc.char.Athletics = v.Athletics
	cc.char.Combat = v.Combat
	cc.char.Engineering = v.Engineering
	cc.char.Pilot = v.Pilot
	cc.char.Science = v.Science
	cc.stage++
	return nil
}

func (cc *CharacterCreator) SelectSpecialAbility(v character.SpecialAbility) error {
	if cc.stage != assignSpecialAbilityStage {
		return fmt.Errorf("can't select special ability in stage %s", cc.stage)
	}
	cc.char.SpecialAbilities = []character.SpecialAbility{v}
	cc.stage++
	return nil
}

func (cc *CharacterCreator) SelectMedKit() error {
	if cc.stage != selectMedKitOrToolkit {
		return fmt.Errorf("can't select a kit in stage %s", cc.stage)
	}
	cc.aquireBasicGear(true)
	return nil
}

func (cc *CharacterCreator) SelectToolKit() error {
	if cc.stage != selectMedKitOrToolkit {
		return fmt.Errorf("can't select a kit in stage %s", cc.stage)
	}
	cc.aquireBasicGear(false)
	return nil
}

func (cc *CharacterCreator) aquireBasicGear(medkit bool) {
	if cc.char.Species.Armor {
		cc.char.Equipment = append(cc.char.Equipment, character.ArmorGear)
	}
	cc.char.Equipment = append(cc.char.Equipment, character.BlasterGear)
	if medkit {
		cc.char.Equipment = append(cc.char.Equipment, character.MedKitGear)
	} else {
		cc.char.Equipment = append(cc.char.Equipment, character.ToolkitGear)
	}
	cc.char.Credits = 500
	cc.stage++
}
