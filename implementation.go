package Lab1

import (
	"fmt"
	"strconv"
	"strings"
)

type Symbol int

const (
	Else Symbol = iota
	Space
	Number
	SimpOperator
	CompOperator
	PowOperator
)

func CharType(c string) Symbol {
	if _, err := strconv.Atoi(c); err == nil {
		return Number
	} else if strings.ContainsAny(c, "+-") {
		return SimpOperator
	} else if strings.ContainsAny(c, "*/") {
		return CompOperator
	} else if c == "^" {
		return PowOperator
	} else if c == " " {
		return Space
	}
	return Else
}

type Prefix struct {
	str []string
	i   int
}

func (p *Prefix) Add(str string) {
	p.str = append(p.str, str)
}

func (p Prefix) Get() string {
	return p.str[p.i]
}
func (p *Prefix) Next() {
	if len(p.str)-1 > p.i {
		p.i++
	}
}

func StringToArray(pref *Prefix, input string) error {
	if len(input) == 0 {
		return fmt.Errorf("Empty argument")
	}
	var res string
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		switch CharType(char) {
		case SimpOperator, CompOperator, PowOperator:
			if res != "" {
				return fmt.Errorf("Invalid syntax")
			}
			res = char
		case Number:
			if res != "" && CharType(res) != Number {
				return fmt.Errorf("Invalid syntax")
			}
			res += char
		case Space:
			if res == "" {
				return fmt.Errorf("Invalid syntax")
			}
			pref.Add(res)
			res = ""
		case Else:
			return fmt.Errorf("Invalid characters")
		}
	}
	if CharType(res) != Number {
		return fmt.Errorf("Invalid syntax")
	}
	pref.Add(res)
	pref.i = 0
	return nil
}

func PrefixToInfix(input string) (string, error) {
	var prefix Prefix
	if err := StringToArray(&prefix, input); err != nil {
		return "", err
	}
	return parse(&prefix, Else), nil
}

func parse(pref *Prefix, prevTSymbol Symbol) string {
	char := pref.Get()
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
	}
	return res
}
