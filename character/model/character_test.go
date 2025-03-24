package model

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

//go:embed sample01.yaml
var sampleChar01 []byte

//go:embed sample02.yaml
var sampleChar02 []byte

//go:embed sample03.yaml
var sampleChar03 []byte

//go:embed sample04.yaml
var sampleChar04 []byte

func TestLoadCharacter(t *testing.T) {

	type tCase struct {
		id                 string
		data               []byte
		expectedSANames    []string
		expectedSACount    int
		expectedIssueIDs   map[string]struct{}
		expectedIssueCount int
	}

	tCases := []tCase{
		{
			id:                 "tc01",
			data:               sampleChar01,
			expectedSACount:    2,
			expectedSANames:    []string{"Forethinker", "Blink*"},
			expectedIssueCount: 0,
			expectedIssueIDs:   map[string]struct{}{},
		},
		{
			id:                 "tc02",
			data:               sampleChar02,
			expectedSACount:    2,
			expectedSANames:    []string{"Forethinker", "Mr. Fixit"},
			expectedIssueCount: 0,
			expectedIssueIDs:   map[string]struct{}{},
		},
		{
			id:                 "tc03",
			data:               sampleChar03,
			expectedSACount:    3,
			expectedSANames:    []string{"Forethinker", "Mr. Fixit", "Resourceful"},
			expectedIssueCount: 0,
			expectedIssueIDs:   map[string]struct{}{},
		},
		{
			id:                 "tc04",
			data:               sampleChar04,
			expectedSACount:    8,
			expectedSANames:    []string{"Hacker", "Resourceful", "Unconventional", "Teleporter Specialist", "Healer", "Forethinker", "Mobile", "Blink (Origins '21)*"},
			expectedIssueCount: 2,
			expectedIssueIDs:   map[string]struct{}{"v01": struct{}{}, "v03": struct{}{}},
		},
	}

	v := NewCharacterValidator()
	for id, tc := range tCases {
		name := fmt.Sprintf("tc%02d", id+1)
		t.Run(name, func(t *testing.T) {
			var c Character
			c, e := LoadCharacter(tc.data)
			if e != nil {
				t.Errorf("error unmarshaling character: %s", e)
				return
			}
			issues, errs := v.ValidateAll(c)
			for _, e := range errs {
				t.Errorf("validation Error: %s", e)
			}
			if len(issues) != tc.expectedIssueCount {
				t.Errorf("expexted %d validation issues but found %d issues", tc.expectedIssueCount, len(issues))
				for _, issue := range issues {
					t.Errorf("unexpexted validation issues: %s", issue)
				}
			}
			for _, issue := range issues {
				if _, ok := tc.expectedIssueIDs[issue.ID()]; !ok {
					t.Errorf("unexpexted validation issues: %s", issue)
				}
			}
			for id := range tc.expectedIssueIDs {
				var idFound bool
				for _, issue := range issues {
					if strings.HasPrefix(issue.String(), id) {
						idFound = true
						break
					}
				}
				if !idFound {
					t.Errorf("failked to find expexted validation issue type %s", id)
				}
			}
			// b, e := json.MarshalIndent(c, "", " ")
			// b, e := json.Marshal(c)
			// if e != nil {
			// 	t.Errorf("error json marshaling character: %s", e)
			// 	return
			// }
			// t.Logf("%s", string(b))
			if len(c.SpecialAbilities) != tc.expectedSACount {
				t.Errorf("error loading special abilities expeted %d, found %d", tc.expectedSACount, len(c.SpecialAbilities))
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
