package character

import (
	_ "embed"
	"testing"

	"github.com/dan-frohlich/battlestations/character/model"
)

//go:embed model/sample.yaml
var sampleChar []byte

func TestManagerLoadAndPrint(t *testing.T) {

	c, err := model.LoadCharacter(sampleChar)
	if err != nil {
		t.Errorf("error loading charatcer: %s", err)
		return
	}

	m := &Manager{character: c}

	m.Print()

}
