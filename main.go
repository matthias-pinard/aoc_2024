package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input/3.input
var input string

const (
	MUL_TOKEN  = "mul("
	DO_TOKEN   = "do()"
	DONT_TOKEN = "don't()"
)

type ParserState struct {
	a   string
	sum int64
	do  bool
}

type Parser interface {
	Feed(c byte, state ParserState) (ParserState, Parser)
}

type Token struct {
	token string
}

type A struct {
	val string
}

type B struct {
	val string
}

func (token Token) Feed(c byte, state ParserState) (ParserState, Parser) {
	token.token += string(c)
	t := token.token
	if t == MUL_TOKEN {
		return state, A{}
	} else if t == DO_TOKEN {
		state.do = true
		return state, Token{}
	} else if t == DONT_TOKEN {
		state.do = false
		return state, Token{}
	} else if strings.HasPrefix(MUL_TOKEN, t) || strings.HasPrefix(DO_TOKEN, t) || strings.HasPrefix(DONT_TOKEN, t) {
		return state, token
	} else {
		return state, Token{}
	}
}

func (a A) Feed(c byte, state ParserState) (ParserState, Parser) {
	if c >= '0' && c <= '9' {
		a.val += string(c)
		return state, a
	} else if c == ',' {
		state.a = a.val
		return state, B{}
	} else {
		return state, Token{}
	}
}

func (b B) Feed(c byte, state ParserState) (ParserState, Parser) {
	if c >= '0' && c <= '9' {
		b.val += string(c)
		return state, b
	} else if c == ')' {
		va, err := strconv.ParseInt(state.a, 10, 64)
		if err != nil {
			panic("error parsing a")
		}
		vb, err := strconv.ParseInt(b.val, 10, 64)
		if err != nil {
			panic("error parsing b")
		}
		if state.do {
			state.sum += va * vb
		}
	}
	return state, Token{}
}

func main() {
	s := ParserState{
		a:   "",
		sum: 0,
		do:  true,
	}
	var p Parser
	p = Token{}
	pos := 0
	for pos < len(input) {
		c := input[pos]
		print(string(c))
		s, p = p.Feed(c, s)
		pos += 1
	}
	println(s.sum)
}
