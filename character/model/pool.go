package model

import (
	"fmt"
	"regexp"
	"strconv"
)

type PoolCode string

func (pc PoolCode) Calculate(c Character) int {
	i, e := strconv.Atoi(string(pc))
	if e == nil { //this is an int
		return i
	}
	switch pc {
	case "athletics":
		return int(c.Athletics)
	case "combat":
		return int(c.Combat)
	case "diplomacy":
		return int(c.Diplomacy)
	case "engineering":
		return int(c.Engineering)
	case "pilot":
		return int(c.Pilot)
	case "psionics":
		return int(c.Psionics)
	case "science":
		return int(c.Science)
	case "sanity":
		return int(c.Sanity)
	default:
		//tokenize the code, then execute the instructions
	}
	return 0
}

type token string

const (
	athToken = "athletics"
	comToken = "combat"
	dipToken = "diplomacy"
	engToken = "engineering"
	pilToken = "pilot"
	psiToken = "psionics"
	sciToken = "science"
	sanToken = "sanity"
	lukToken = "luck"
	rnkToken = "rank"
)

func checkToken(runes []rune) (token, bool) {
	switch string(runes) {
	case "x", "*", "+", athToken, comToken, dipToken, engToken, pilToken, psiToken, sciToken, sanToken, lukToken, rnkToken:
		return token(runes), true
	}
	return token(""), false
}

func tokenizer(pc PoolCode) (tokens []token, err error) {
	runes := []rune(pc)
	max := len(runes)
	for i := 0; i < max; i++ {

		r := runes[i]
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j := i
			numRunes := []rune("0123456789")
			runeSet := make(map[rune]struct{})
			for _, r := range string(numRunes) {
				runeSet[r] = struct{}{}
			}
			reg := regexp.MustCompile("[0-9][0-9]*")
			_, isNum := runeSet[runes[j]]
			for isNum {
				j++
				if j == max {
					break
				}
				_, isNum = runeSet[runes[j]]
			}
			if !reg.Match([]byte(string(runes[i:j]))) {
				return tokens, fmt.Errorf("expected %s to be a number", string(runes[i:j]))
			}

			tokens = append(tokens, token(runes[i:j]))
			i = j - 1

		case 'x', '*', '+':
			tokens = append(tokens, token(runes[i]))
		case 'a':
			if len(runes[i:]) >= len(athToken) {
				if t, ok := checkToken(runes[i : len(athToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(athToken) - 1
				}
			}
		case 'c':
			if len(runes[i:]) >= len(comToken) {
				if t, ok := checkToken(runes[i : len(comToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(comToken) - 1
				}
			}
		case 'd':
			if len(runes[i:]) >= len(dipToken) {
				if t, ok := checkToken(runes[i : len(dipToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(dipToken) - 1
				}
			}
		case 'e':
			if len(runes[i:]) >= len(engToken) {
				if t, ok := checkToken(runes[i : len(engToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(engToken) - 1
				}
			}
		case 'l':
			if len(runes[i:]) >= len(lukToken) {
				if t, ok := checkToken(runes[i : len(lukToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(lukToken) - 1
				}
			}
		case 'p':
			if len(runes[i:]) >= len(pilToken) {
				if t, ok := checkToken(runes[i : len(pilToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(pilToken) - 1
				}
			}
			if len(runes[i:]) >= len(psiToken) {
				if t, ok := checkToken(runes[i : len(psiToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(psiToken) - 1
				}
			}
		case 'r':
			if len(runes[i:]) >= len(rnkToken) {
				if t, ok := checkToken(runes[i : len(rnkToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(rnkToken) - 1
				}
			}
		case 's':
			if len(runes[i:]) >= len(sciToken) {
				if t, ok := checkToken(runes[i : len(sciToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(sciToken) - 1
				}
			}
			if len(runes[i:]) >= len(sanToken) {
				if t, ok := checkToken(runes[i : len(sanToken)+i]); ok {
					tokens = append(tokens, t)
					i += len(sanToken) - 1
				}
			}
		default:
		}
	}
	return tokens, err
}
