package Lab1

import (
	"strings"
)

type Symbol int

const (
	Else Symbol = iota
	Number
	SimpOperator
	CompOperator
	PowOperator
)

func CharType(c string) Symbol {
	if c >= "0" && c <= "9" {
		return Number
	} else if strings.ContainsAny(c, "+-") {
		return SimpOperator
	} else if strings.ContainsAny(c, "*/") {
		return CompOperator
	} else if strings.ContainsAny(c, "^") {
		return PowOperator
	}
	return Else
}

type Prefix struct {
	str string
	i   int
}

func (p Prefix) Current() byte {
	return p.str[p.i]
}
func (p *Prefix) Next() {
	if len(p.str)-1 > p.i {
		p.i++
	}
}

func PrefixToInfix(input string) (string, error) {
	return parse(&Prefix{input, 0}, Else), nil
}

func parse(pref *Prefix, prevTSymbol Symbol) string {
	char := string(pref.Current())
	pref.Next()
	tSymbol := CharType(char)

	var res string
	switch tSymbol {
	case SimpOperator, CompOperator, PowOperator:
		res = parse(pref, tSymbol) + char + parse(pref, tSymbol)
		if tSymbol < prevTSymbol {
			res = "(" + res + ")"
		}
	case Number:
		res = char
	case Else:
		res = parse(pref, prevTSymbol)
	}
	return res
}
