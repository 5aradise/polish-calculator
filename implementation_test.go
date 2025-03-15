package poca

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	res, err := PostfixToLisp("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(+ 5 (* (- 4 2) (pow 3 2)))", res)
	}
}

func ExamplePostfixToLisp() {
	res, _ := PostfixToLisp("2 2 +")
	fmt.Println(res)

	// Output:
	// (+ 2 2)
}
