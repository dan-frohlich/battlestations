package character

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/dan-frohlich/battlestations/character/model"
)

var (
	//go:embed model/sample01.yaml
	sampleChar01 []byte

	//go:embed model/sample02.yaml
	sampleChar02 []byte

	//go:embed model/sample03.yaml
	sampleChar03 []byte

	//go:embed model/sample04.yaml
	sampleChar04 []byte
)

func TestManagerLoadAndPrint(t *testing.T) {

	tCases := [][]byte{sampleChar01, sampleChar02, sampleChar03, sampleChar04}

	for id, tc := range tCases {
		t.Run(fmt.Sprintf("tc:%02d", id+1),
			func(t *testing.T) {
				c, err := model.LoadCharacter(tc)
				if err != nil {
					t.Errorf("error loading charatcer: %s", err)
					return
				}
				m := &Manager{character: c}
				m.Print()
			})
	}
}
