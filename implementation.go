package poca

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidInput     = errors.New("poca: invalid input")
	ErrWrongNumOfTokens = errors.New("wrong number of tokens")
	ErrBadToken         = errors.New("bad token")
	errInvalidNodeType  = errors.New("invalid node type (must be *expNode or int)")
)

// PostfixToLisp converts a mathematical expression written in postfix notation
// (Reverse Polish Notation, RPN) into a Lisp-style expression.
//
// If the input string is empty or contains an invalid expression, an error is returned.
func PostfixToLisp(input string) (string, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", fmt.Errorf("%w: empty value", ErrInvalidInput)
	}

	tree, num, err := expTreeOrNumFromPostfix(tokens)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInvalidInput, err)
	}

	if tree == nil {
		return strconv.Itoa(num), nil
	}

	return tree.toLisp()
}

var polishToLispOp = map[string]string{
	"+": "+",
	"-": "-",
	"*": "*",
	"/": "/",
	"^": "pow",
}

type expNode struct {
	op    string
	left  any
	right any
}

func expTreeOrNumFromPostfix(tokens []string) (*expNode, int, error) {
	var nodes stack[any]

	for _, token := range tokens {
		if _, ok := polishToLispOp[token]; ok {
			right, ok := nodes.pop()
			if !ok {
				return nil, 0, ErrWrongNumOfTokens
			}
			left, ok := nodes.pop()
			if !ok {
				return nil, 0, ErrWrongNumOfTokens
			}
			nodes.push(&expNode{token, left, right})
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return nil, 0, fmt.Errorf("%w: %w", ErrBadToken, err)
			}
			nodes.push(num)
		}
	}

	node, ok := nodes.pop()
	if !ok {
		return nil, 0, ErrWrongNumOfTokens
	}
	_, ok = nodes.pop()
	if ok {
		return nil, 0, ErrWrongNumOfTokens
	}

	switch exp := node.(type) {
	case *expNode:
		return exp, 0, nil
	case int:
		return nil, exp, nil
	default:
		return nil, 0, nil
	}
}

func (tree *expNode) toLisp() (string, error) {
	if tree == nil {
		return "", errors.New("converting nil pointer")
	}

	var b strings.Builder

	b.WriteByte('(')
	b.WriteString(polishToLispOp[tree.op])
	b.WriteByte(' ')
	switch node := tree.left.(type) {
	case *expNode:
		exp, err := node.toLisp()
		if err != nil {
			return "", err
		}
		b.WriteString(exp)
	case int:
		b.WriteString(strconv.Itoa(node))
	default:
		return "", errInvalidNodeType
	}
	b.WriteByte(' ')
	switch node := tree.right.(type) {
	case *expNode:
		exp, err := node.toLisp()
		if err != nil {
			return "", err
		}
		b.WriteString(exp)
	case int:
		b.WriteString(strconv.Itoa(node))
	default:
		return "", errInvalidNodeType
	}
	b.WriteByte(')')

	return b.String(), nil
}
