package poca

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToLisp(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		res  string
		err  error
	}{
		{
			name: "single number",
			arg:  "5",
			res:  "5",
			err:  nil,
		},
		{
			name: "basic addition",
			arg:  "3 4 +",
			res:  "(+ 3 4)",
			err:  nil,
		},
		{
			name: "power conversion",
			arg:  "45 78 0 + ^",
			res:  "(pow 45 (+ 78 0))",
			err:  nil,
		},
		{
			name: "negative value",
			arg:  "28 52 / -2 -",
			res:  "(- (/ 28 52) -2)",
			err:  nil,
		},
		{
			name: "alternation of operations",
			arg:  "69 69 69 69 69 69 69 ^ / - + * ^",
			res:  "(pow 69 (* 69 (+ 69 (- 69 (/ 69 (pow 69 69))))))",
			err:  nil,
		},
		{
			name: "big expression 1",
			arg:  "25 32 5 7 ^ - 44 37 + - *",
			res:  "(* 25 (- (- 32 (pow 5 7)) (+ 44 37)))",
			err:  nil,
		},
		{
			name: "big expression 2",
			arg:  "21 24 39 - 9 88 90 2 - ^ + 3 - / 0 * - 50 45 * +",
			res:  "(+ (- 21 (* (/ (- 24 39) (- (+ 9 (pow 88 (- 90 2))) 3)) 0)) (* 50 45))",
			err:  nil,
		},
		{
			name: "big expression 3",
			arg:  "42 69 + 49 28 99 ^ 1 25 4 ^ / - * 43 100 ^ / - 420 69 + -",
			res:  "(- (- (+ 42 69) (/ (* 49 (- (pow 28 99) (/ 1 (pow 25 4)))) (pow 43 100))) (+ 420 69))",
			err:  nil,
		},
		{
			name: "empty string",
			arg:  "",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "only spaces",
			arg:  "     ",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "only numbers",
			arg:  "47 22 35",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "invalid operator 1",
			arg:  "58 0 div",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "invalid operator 2",
			arg:  "69 69 _",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "wrong number of tokens 1",
			arg:  "69 69 + +",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "wrong number of tokens 2",
			arg:  "69 69 69 69 - / + ^",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "wrong token sequence 1",
			arg:  "69 + 69 - 69 * 69",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "wrong token sequence 2",
			arg:  "+ / 69 69 / 69 69",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "wrong token sequence 3",
			arg:  "69 + 69 69 69 / + ^",
			res:  "",
			err:  ErrInvalidInput,
		},
		{
			name: "wrong token sequence 4",
			arg:  "69 69 - 69 / * 69",
			res:  "",
			err:  ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			res, err := PostfixToLisp(tt.arg)
			if tt.err == nil {
				if assert.Nil(err) {
					assert.Equal(tt.res, res)
				}
			} else {
				assert.ErrorIs(err, tt.err)
			}
		})
	}
}

func ExamplePostfixToLisp() {
	res, _ := PostfixToLisp("5 4 2 - 3 2 ^ * +")
	fmt.Println(res)

	// Output:
	// (+ 5 (* (- 4 2) (pow 3 2)))
}
