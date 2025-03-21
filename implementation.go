package lab2

import (
	"errors"
	"strings"
)

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

// PostfixToInfix перетворює постфіксний вираз у інфіксний
func PostfixToInfix(input string) (string, error) {
	postfix := strings.Fields(input)
	var stack []string

	for _, token := range postfix {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("недостатньо операндів у виразі")
			}

			// Два останні операнди зі стеку
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Формуємо інфіксний вираз і додаємо його назад у стек
			expr := "(" + op1 + " " + token + " " + op2 + ")"
			stack = append(stack, expr)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("некоректний постфіксний вираз")
	}

	return stack[0], nil
}
