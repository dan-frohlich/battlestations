package model

import "testing"

func TestPoolTokenizer(t *testing.T) {
	// c, err := LoadCharacter(sampleChar)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	type tCase struct {
		id                 string
		code               PoolCode
		expectedTokenCount int
		expectedTokens     []token
	}
	tCases := []tCase{
		{
			id:                 "tc01",
			code:               PoolCode("athletics + 2"),
			expectedTokenCount: 3,
			expectedTokens:     []token{"athletics", "+", "2"},
		},
		{
			id:                 "tc01a",
			code:               PoolCode("2 + athletics"),
			expectedTokenCount: 3,
			expectedTokens:     []token{"2", "+", "athletics"},
		},
		{
			id:                 "tc02",
			code:               PoolCode("athletics + psionics"),
			expectedTokenCount: 3,
			expectedTokens:     []token{"athletics", "+", "psionics"},
		},
		{
			id:                 "tc03",
			code:               PoolCode("rankx2"),
			expectedTokenCount: 3,
			expectedTokens:     []token{"rank", "x", "2"},
		},
		{
			id:                 "tc03a",
			code:               PoolCode("2xrank"),
			expectedTokenCount: 3,
			expectedTokens:     []token{"2", "x", "rank"},
		},
		{
			id:                 "tc04",
			code:               PoolCode("5"),
			expectedTokenCount: 1,
			expectedTokens:     []token{"5"},
		},
		{
			id:                 "tc05",
			code:               PoolCode("science"),
			expectedTokenCount: 1,
			expectedTokens:     []token{"science"},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.id, func(t *testing.T) {

			pc := tc.code

			tokens, err := tokenizer(pc)
			if err != nil {
				t.Errorf("unexpected tokenizer error: %s", err)
				return
			}
			expectedTokenCount := tc.expectedTokenCount
			if len(tokens) != expectedTokenCount {
				t.Errorf("expected %d tokens but found %d for [%s]", expectedTokenCount, len(tokens), pc)
			}

			expectedTokens := tc.expectedTokens
			for i := range tokens {
				if i > len(expectedTokens) {
					break
				}
				if tokens[i] != expectedTokens[i] {
					t.Errorf("expected token[%d] to be [%s] but found [%s]", i, expectedTokens[i], tokens[i])
				}
			}
		})
	}
}
