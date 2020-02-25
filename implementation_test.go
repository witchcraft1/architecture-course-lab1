package Lab1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfixSimple(t *testing.T) {
	res, err := PrefixToInfix("+23")
	if assert.Nil(t, err) {
		assert.Equal(t, "2+3", res)
	}
	res, err = PrefixToInfix("*+723")
	if assert.Nil(t, err) {
		assert.Equal(t, "(7+2)*3", res)
	}
	res, err = PrefixToInfix("^7*23")
	if assert.Nil(t, err) {
		assert.Equal(t, "7^(2*3)", res)
	}
}

func TestPrefixToInfixComplex(t *testing.T) {
	res, err := PrefixToInfix("+5-^74*3/^375")
	if assert.Nil(t, err) {
		assert.Equal(t, "5+7^4-3*3^7/5", res)
	}
	res, err = PrefixToInfix("+5/^^7-4*3375")
	if assert.Nil(t, err) {
		assert.Equal(t, "5+7^(4-3*3)^7/5", res)
	}
}

func TestPrefixToInfixEmpty(t *testing.T) {
	res, err := PrefixToInfix("   ")
	if assert.Nil(t, err) {
		assert.Equal(t, "", res)
	}
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2+2
}
