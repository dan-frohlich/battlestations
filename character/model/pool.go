package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PoolCode string

func (pc PoolCode) Calculate(c Character) (i int, e error) {
	i, e = strconv.Atoi(string(pc))
	if e == nil { //this is an int
		return i, e
	}
	switch pc {
	case "athletics":
		return int(c.Athletics), nil
	case "combat":
		return int(c.Combat), nil
	case "diplomacy":
		return int(c.Diplomacy), nil
	case "engineering":
		return int(c.Engineering), nil
	case "pilot":
		return int(c.Pilot), nil
	case "psionics":
		return int(c.Psionics), nil
	case "science":
		return int(c.Science), nil
	case "sanity":
		return int(c.Sanity), nil
	default:
		//tokenize the code, then execute the instructions
		var (
			tokens []token
			ast    *node
		)
		tokens, e = tokenizer(pc)
		if e != nil { //failed to tokenize
			return i, e
		}
		ast, e = lexer(tokens)
		if e != nil { //failed to lex
			return i, e
		}
		v := ast.eval(c)
		return v, nil
	}
}

type token string

func (t token) isOperator() bool {
	switch t {
	case "x", "+", "*", "_":
		return true
	}
	return false
}

var isNumericRegexp = regexp.MustCompile("[0-9][0-9]*")

func (t token) isNumeric() bool {
	return isNumericRegexp.Match([]byte(t))
}

func (t token) evalAsCharacterReference(c Character) int {
	switch t {
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
	}
	return -1
}

func (t token) asNumeric() int {
	if !t.isNumeric() {
		return -1
	}
	i, e := strconv.Atoi(string(t))
	if e != nil {
		return -1
	}
	return i
}

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
	case "x", "*", "+", "_", athToken, comToken, dipToken, engToken, pilToken, psiToken, sciToken, sanToken, lukToken, rnkToken:
		return token(runes), true
	}
	return token(""), false
}

func tokenizer(pc PoolCode) (tokens []token, err error) {
	runes := []rune(strings.ToLower(string(pc)))
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
			_, isNum := runeSet[runes[j]]
			for isNum {
				j++
				if j == max {
					break
				}
				_, isNum = runeSet[runes[j]]
			}
			if !isNumericRegexp.Match([]byte(string(runes[i:j]))) {
				return tokens, fmt.Errorf("expected %s to be a number", string(runes[i:j]))
			}

			tokens = append(tokens, token(runes[i:j]))
			i = j - 1

		case 'x', '*', '+', '_':
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

type node struct {
	raw   token
	value int
	left  *node
	right *node
}

func (ast *node) eval(c Character) int {
	n := ast
	switch {
	case n.raw.isNumeric():
		return n.value
	case n.raw.isOperator():
		switch n.raw {
		case "+":
			return n.left.eval(c) + n.right.eval(c)
		case "x", "*":
			return n.left.eval(c) * n.right.eval(c)
		case "_":
			return min(n.left.eval(c), n.right.eval(c))
		default:
			return -1
		}
	default:
		return n.raw.evalAsCharacterReference(c)
	}
}

func min(arg int, args ...int) int {
	m := arg
	for _, a := range args {
		if a < m {
			m = a
		}
	}
	return m
}

func lexer(tokens []token) (ast *node, err error) {
	var head *node

	var first = true
	var expectValue = true
	for _, t := range tokens {
		if first {
			first = false
			//first token must ve a value
			if t.isOperator() {
				return nil, fmt.Errorf("can't start pool code with an operator: %s", t)
			}
		}
		if expectValue && t.isOperator() {
			return nil, fmt.Errorf("expected a value but found an operator: %s", t)
		}
		if !expectValue && !t.isOperator() {
			return nil, fmt.Errorf("expected a operator but found a value: %s", t)
		}
		switch {
		case t.isOperator():
			n := &node{
				raw:  t,
				left: head,
			}
			head = n
		default:
			n := &node{
				raw: t,
			}
			if t.isNumeric() {
				n.value = t.asNumeric()
			}
			switch {
			case head == nil:
				head = n
			case head.left == nil:
				head.left = n
			default:
				head.right = n
			}
		}
		expectValue = !expectValue
	}
	return head, nil
}
