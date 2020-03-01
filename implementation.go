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
	if _, err := strconv.Atoi(c); err == nil && !strings.ContainsAny(c, "+-") {
		return Number
	} else if c == "+" || c == "-" {
		return SimpOperator
	} else if c == "*" || c == "/" {
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

func (p *Prefix) Get() (string, error) {
	//check if too many operators
	if len(p.str) <= p.i {
		return "", fmt.Errorf("Invalid Syntax")
	}
	res := p.str[p.i]
	p.i++
	return res, nil
}

//This function converts prefix expression(+ * 2 2 * - 10 2 5) to infix expression(2*2+(10-2)*5).
//To convert postfix to infix you have to pass expression as a string.
//Possible results are infix form of expression or error if it occurs.
func PrefixToInfix(input string) (string, error) {
	var prefix Prefix
	prefix.str = strings.Split(input, " ")
	//check expression and build new expression through recursive function
	res, err := parse(&prefix, Else)
	//check if too many operands
	if prefix.i != len(prefix.str) {
		return res, fmt.Errorf("Too many operators or operands")
	}
	return res, err
}

func parse(pref *Prefix, prevTSymbol Symbol) (string, error) {
	char, err := pref.Get()
	tSymbol := CharType(char)

	var res string
	switch tSymbol {
	case SimpOperator, CompOperator, PowOperator:
		op1, err1 := parse(pref, tSymbol)
		op2, err2 := parse(pref, tSymbol)
		if err1 != nil {
			err = err1
		} else if err2 != nil {
			err = err2
		}
		res = op1 + char + op2
		if tSymbol < prevTSymbol {
			res = "(" + res + ")"
		}
	case Number:
		res = char
	case Else:
		err = fmt.Errorf("Invalid Syntax")
	}
	return res, err
}
