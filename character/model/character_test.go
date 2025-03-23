package model

import (
	_ "embed"
	"testing"

	"gopkg.in/yaml.v2"
)

//go:embed sample.yaml
var sampleChar []byte

func TestLoadCharacter(t *testing.T) {

	var c Character
	c, e := LoadCharacter(sampleChar)
	if e != nil {
		t.Errorf("error unmarcshaling character: %s", e)
		return
	}
	t.Logf("%#v", c)
	if len(c.SpecialAbilities) != 1 {
		t.Errorf("error loading special abilities: %d", len(c.SpecialAbilities))
		return
	}
	if c.SpecialAbilities[0].Name != "Forethinker" {
		t.Errorf("error loading Forethinker: %s", c.SpecialAbilities[0].Name)
	}
	if c.SpecialAbilities[0].Summary == "" {
		t.Errorf("error loading Forethinker.Summary: %s", c.SpecialAbilities[0].Summary)
	}
}

func TestSA(t *testing.T) {
	const tc = "Forethinker"

	var failed bool
	a := GetAbility(tc)
	if a.Name != tc {
		t.Errorf("did not find %s", tc)
		failed = true
	}
	if a.Summary == "" {
		t.Errorf("did not find %s.Summary", tc)
		failed = true
	}

	if failed {
		b, e := yaml.Marshal(abilityIndex)
		if e != nil {
			t.Errorf("did not marshal ability registry: %s", e)
			return
		}
		t.Log(string(b))
	}
}
