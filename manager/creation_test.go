package manager

import (
	"testing"

	"github.com/dan-frohlich/battlestations/character"
	"gopkg.in/yaml.v2"
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
	err = cc.SelectSpecies(character.CanosianSpecies)
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

	t.Logf("char: %#v", cc.char)

	bsc := convertForPrinting(cc.char)

	var b []byte
	b, err = yaml.Marshal(bsc)

	t.Logf("\n%s", b)
	// ioutil.WriteFile("test_char.yml", b, 0666)
}
