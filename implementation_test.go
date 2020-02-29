package Lab1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfixSimple(t *testing.T) {
	res, err := PrefixToInfix("+ 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "2+3", res)
	}
	res, err = PrefixToInfix("* + 7 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "(7+2)*3", res)
	}
	res, err = PrefixToInfix("^ 7 * 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "7^(2*3)", res)
	}
	res, err = PrefixToInfix("* 7 23")
	if assert.Nil(t, err) {
		assert.Equal(t, "7*23", res)
	}
}

func TestPrefixToInfixComplex(t *testing.T) {
	res, err := PrefixToInfix("+ 5 - ^ 7 4 * 3 / ^ 3 7 5")
	if assert.Nil(t, err) {
		assert.Equal(t, "5+7^4-3*3^7/5", res)
	}
	res, err = PrefixToInfix("+ 5 / ^ ^ 7 - 4 * 3 3 7 5")
	if assert.Nil(t, err) {
		assert.Equal(t, "5+7^(4-3*3)^7/5", res)
	}
	res, err = PrefixToInfix("- * ^ + 2 13 / + 27 653 340 3 / 560 ^ - 18 4 10")
	if assert.Nil(t, err) {
		assert.Equal(t, "(2+13)^((27+653)/340)*3-560/(18-4)^10", res)
	}
}

func TestPrefixToInfixError(t *testing.T) {
	_, err := PrefixToInfix("   ")
	assert.Error(t, err)
	_, err = PrefixToInfix("+2155")
	assert.Error(t, err)
	_, err = PrefixToInfix("+21* 55 52")
	assert.Error(t, err)
	_, err = PrefixToInfix("*- 122")
	assert.Error(t, err)
	_, err = PrefixToInfix("- 125 *- 122 144 10")
	assert.Error(t, err)
	_, err = PrefixToInfix("+ * 7 9 6 2 4")
	assert.Error(t, err)
	_, err = PrefixToInfix("+ * 2 4")
	assert.Error(t, err)
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2+2
}
