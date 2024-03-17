package manager

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/dan-frohlich/battlestations/character"
	"github.com/dan-frohlich/battlestations/print"
)

func TestCreateCharacter(t *testing.T) {
	cc := NewCharacterCreator()

	cc.SetName("TestCreateCharacter")

	t.Logf("stage: %s", cc.stage)

	err := cc.SelectSpecies(character.CanosianSpecies)
	t.Logf("stage: %s", cc.stage)

	if err != nil {
		t.Errorf("FAIL: SelectSpecies, %s", err)
	}

	err = cc.SelectSpecies(character.CanosianSpecies)
	t.Logf("stage: %s", cc.stage)

	if err == nil {
		t.Errorf("FAIL: SelectSpecies should not be possible after we select a species")
		return
	}

	cc.PreviousStage()
	t.Logf("stage: %s", cc.stage)
	err = cc.SelectSpecies(character.XeloxianSpecies)
	t.Logf("stage: %s", cc.stage)

	if err != nil {
		t.Errorf("FAIL: SelectSpecies, %s", err)
		return
	}

	err = cc.SelectProfession(character.PilotProfession)
	t.Logf("stage: %s", cc.stage)

	if err != nil {
		t.Errorf("FAIL: SelectProfession, %s", err)
		return
	}

	cc.PreviousStage()
	t.Logf("stage: %s", cc.stage)
	err = cc.SelectProfession(character.ScientistProfession)
	t.Logf("stage: %s", cc.stage)
	if err != nil {
		t.Errorf("FAIL: SelectProfession, %s", err)
		return
	}

	err = cc.SelectSkills(SkillSelection{
		Athletics:   0,
		Combat:      0,
		Engineering: 0,
		Pilot:       0,
		Science:     0,
	})

	t.Logf("stage: %s", cc.stage)
	if err == nil {
		t.Errorf("FAIL: SelectSkills should not be possible with all 0 skills")
		return
	}

	levels := cc.GetSkillArrays()[3]

	err = cc.SelectSkills(SkillSelection{
		Athletics:   levels[0],
		Combat:      levels[0],
		Engineering: levels[0],
		Pilot:       levels[0],
		Science:     levels[0],
	})
	t.Logf("stage: %s", cc.stage)

	if err == nil {
		t.Errorf("FAIL: SelectSkills should not be possible with all maxed skills")
		return
	}

	err = cc.SelectSkills(SkillSelection{
		Athletics:   levels[3],
		Combat:      levels[2],
		Engineering: levels[1],
		Pilot:       levels[4],
		Science:     levels[0],
	})
	t.Logf("stage: %s", cc.stage)

	if err != nil {
		t.Errorf("FAIL: SelectProfession, %s", err)
		return
	}

	err = cc.SelectSpecialAbility(character.BosunSpecialAbility)
	t.Logf("stage: %s", cc.stage)

	if err != nil {
		t.Errorf("FAIL: SelectSpecialAbility, %s", err)
		return
	}

	err = cc.SelectMedKit()
	t.Logf("stage: %s", cc.stage)

	if err != nil {
		t.Errorf("FAIL: SelectSpecialAbility, %s", err)
		return
	}

	t.Logf("char: %#v", cc.char)

	bsc := convertForPrinting(cc.char)

	err = print.WritePDFFile(bsc)
	if err != nil {
		t.Errorf("FAIL: WritePDFFile, %s", err)
		return
	}
	var b bytes.Buffer
	foo := bufio.NewWriter(&b)

	err = print.WritePDF(bsc, foo) //TODO which write closer?
	if err != nil {
		t.Errorf("FAIL: WritePDF, %s", err)
		return
	}

	t.Logf("created pdf of %d bytes", len(b.Bytes()))
	t.Log(hex.EncodeToString(b.Bytes())[:64], "...")

}
