package model

import (
	_ "embed"
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v2"
)

//go:embed sample.yaml
var sampleChar []byte

func TestLoadCharacter(t *testing.T) {

	var c Character
	c, e := LoadCharacter(sampleChar)
	if e != nil {
		t.Errorf("error unmarshaling character: %s", e)
		return
	}
	// b, e := json.MarshalIndent(c, "", " ")
	b, e := json.Marshal(c)
	if e != nil {
		t.Errorf("error json marshaling character: %s", e)
		return
	}
	t.Logf("%s", string(b))
	if len(c.SpecialAbilities) != 2 {
		t.Errorf("error loading special abilities: %d", len(c.SpecialAbilities))
		return
	}
	if c.SpecialAbilities[0].Name != "Forethinker" {
		t.Errorf("error loading Forethinker: %s", c.SpecialAbilities[0].Name)
	}
	if c.SpecialAbilities[0].Summary == "" {
		t.Errorf("error loading Forethinker.Summary: %s", c.SpecialAbilities[0].Summary)
	}
	if c.SpecialAbilities[1].Name != "Blink"+ItemNotFoundString {
		t.Errorf("error loading Blink: %s", c.SpecialAbilities[1].Name)
	}
	if c.SpecialAbilities[1].Summary == "" {
		t.Errorf("error loading Blink.Summary: %s", c.SpecialAbilities[1].Summary)
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
