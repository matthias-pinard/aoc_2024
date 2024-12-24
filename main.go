package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input/3.input
var input string

const (
	MUL_TOKEN  = "mul"
	DO_TOKEN   = "do()"
	DONT_TOKEN = "don't()"
)

type ParserState uint32

const (
	TOKEN   ParserState = 0
	LEFT_P  ParserState = 1
	A       ParserState = 2
	COMMA   ParserState = 3
	B       ParserState = 4
	RIGHT_P ParserState = 5
	MUL     ParserState = 6
	DO      ParserState = 7
	DONT    ParserState = 8
)

func main() {
	var r int64
	a := ""
	b := ""
	pos := 0

	s := TOKEN

	t := ""
	d := true

	for pos != len(input) {
		c := input[pos]
		pos += 1
		if s == TOKEN {
			t += string(c)
			if t == MUL_TOKEN {
				t = ""
				s = LEFT_P
			} else if t == DO_TOKEN {
				d = true
				t = ""
				s = TOKEN
			} else if t == DONT_TOKEN {
				d = false
				t = ""
				s = TOKEN
			} else if strings.HasPrefix(MUL_TOKEN, t) || strings.HasPrefix(DO_TOKEN, t) || strings.HasPrefix(DONT_TOKEN, t) {
				// Do nothings
			} else {
				t = ""
				s = TOKEN
			}
			continue
		}

		if s == LEFT_P {
			if c == '(' {
				s = A
			} else {
				s = TOKEN
			}
			continue
		}

		if s == A {
			if c >= '0' && c <= '9' {
				a += string(c)
			} else if c == ',' {
				s = B
			} else {
				a = ""
				s = TOKEN
			}
			continue
		}

		if s == B {
			if c >= '0' && c <= '9' {
				b += string(c)
			} else if c == ')' {
				va, err := strconv.ParseInt(a, 10, 64)
				if err != nil {
					panic("error parsing a")
				}
				vb, err := strconv.ParseInt(b, 10, 64)
				if err != nil {
					panic("error parsing b")
				}
				if d {
					r += va * vb
				}
				a = ""
				b = ""
				s = TOKEN
			} else {
				a = ""
				b = ""
				s = TOKEN
			}
			continue
		}
	}
	println(r)
}
