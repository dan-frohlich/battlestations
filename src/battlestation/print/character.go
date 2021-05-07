package main

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type bsChar map[string]interface{}

func loadCharFromReader( r io.Reader) bsChar {
	t := bsChar{}
	data, e := ioutil.ReadAll(r)
	check(e, "read char")
	e = yaml.Unmarshal(data, &t)
	check(e, "unmarshal  char")
	return t
}

func (charData bsChar) isLarge() bool {
	if val, ok := charData["special_abilities"]; ok {
		if sas, ok := val.([]interface{}); ok {
			if len(sas) > 10 {
				return true
			}
		}
	}

	if val, ok := charData["equipment"]; ok {
		if equip, ok := val.([]interface{}); ok {
			if len(equip) > 10 {
				return true
			}
		}
	}
	return false
}
