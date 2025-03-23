package model

import (
	"strings"
	"testing"
)

func TestDBPoolCodes(t *testing.T) {

	var c Character
	for _, sa := range abilities {
		if sa.Pool != "" {
			_, e := sa.Pool.Calculate(c)
			if e != nil {
				t.Errorf("bad pool code: [%s]: %s", sa.Pool, e)
				// } else {
				// t.Logf("OK pool validated: [%s]", sa.Pool)
			}
		} else {
			if strings.Contains(strings.ToLower(sa.Summary), " pool of ") {
				t.Errorf("no pool code for: [%s]: %s", sa.Name, sa.Summary)
			}
		}
	}
}
