package lab2

import (
	"errors"
	"strconv"
	"strings"
)

var validOperators = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
	"^": true,
}

func PostfixToInfix(expression string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			stack = append(stack, token)
		} else if validOperators[token] {
			if len(stack) < 2 {
				return "", errors.New("invalid expression: not enough operands")
			}

			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			newExpr := "(" + operand1 + " " + token + " " + operand2 + ")"
			stack = append(stack, newExpr)
		} else {
			return "", errors.New("invalid token: " + token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid expression: too many operands")
	}

	return stack[0], nil
}
