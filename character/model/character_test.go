package model

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

//go:embed sample01.yaml
var sampleChar01 []byte

//go:embed sample02.yaml
var sampleChar02 []byte

func TestLoadCharacter(t *testing.T) {

	type tCase struct {
		id              string
		data            []byte
		expectedSANames []string
		expectedSACount int
	}

	tCases := []tCase{
		{
			id:              "tc01",
			data:            sampleChar01,
			expectedSACount: 2,
			expectedSANames: []string{"Forethinker", "Blink*"},
		},
		{
			id:              "tc02",
			data:            sampleChar02,
			expectedSACount: 2,
			expectedSANames: []string{"Forethinker", "Mr. Fixit"},
		},
	}

	for id, tc := range tCases {
		name := fmt.Sprintf("tc%02d", id+1)
		t.Run(name, func(t *testing.T) {
			var c Character
			c, e := LoadCharacter(tc.data)
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
			if len(c.SpecialAbilities) != tc.expectedSACount {
				t.Errorf("error loading special abilities expeted %d, found %d", tc.expectedSACount, len(c.SpecialAbilities))
				return
			}
			for i := range c.SpecialAbilities {
				if i >= len(tc.expectedSANames) {
					t.Errorf("unexpected Special Ability: %s", c.SpecialAbilities[i].Name)
					continue
				}
				if c.SpecialAbilities[i].Name != tc.expectedSANames[i] {
					t.Errorf("expected special ability %d to be %s but found: %s",
						i, tc.expectedSANames[i], c.SpecialAbilities[i].Name)
				}
			}
		})
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
