package lab2

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	tests := []struct {
		name        string
		input  string
		expected    string
		err bool
	}{
		{
			name:        "Simple addition",
			input:  	 "3 4 +",
			expected:    "(3 + 4)",
			err: false,
		},
		{
			name:        "Simple subtraction",
			input:  	 "10 5 -",
			expected:    "(10 - 5)",
			err: false,
		},
		{
			name:        "Simple multiplication",
			input:  	 "6 7 *",
			expected:    "(6 * 7)",
			err: false,
		},
		{
			name:        "Simple division",
			input:  	 "20 4 /",
			expected:    "(20 / 4)",
			err: false,
		},
		{
			name:        "Simple exponentiation",
			input:  	 "2 3 ^",
			expected:    "(2 ^ 3)",
			err: false,
		},
		{
			name:        "Simple expression with 3 operands",
			input:  	 "3 4 + 5 *",
			expected:    "((3 + 4) * 5)",
			err: false,
		},
		{
			name:        "Simple expression with 3 operands (mixed operators)",
			input:  	 "10 2 / 5 +",
			expected:    "((10 / 2) + 5)",
			err: false,
		},
		{
			name:        "Complex expression 1",
			input:  	 "2 3 ^ 4 5 + * 6 - 7 /",
			expected:    "((((2 ^ 3) * (4 + 5)) - 6) / 7)",
			err: false,
		},
		{
			name:        "Complex expression 2",
			input:  	 "1 2 + 3 * 4 5 ^ - 6 / 7 +",
			expected:    "(((((1 + 2) * 3) - (4 ^ 5)) / 6) + 7)",
			err: false,
		},
		{
			name:        "Complex expression 3",
			input:  	 "10 2 / 3 * 4 2 ^ + 5 - 6 7 + *",
			expected:    "(((((10 / 2) * 3) + (4 ^ 2)) - 5) * (6 + 7))",
			err: false,
		},
		{
			name:        "Complex expression 4",
			input:  	 "15 3 / 5 + 2 4 * - 10 + 6 2 ^ / 3 *",
			expected:    "((((((15 / 3) + 5) - (2 * 4)) + 10) / (6 ^ 2)) * 3)",
			err: false,
		},
		{
			name:        "Complex expression 5",
			input:  	 "2 3 ^ 4 5 + * 6 - 7 / 8 9 + *",
			expected:    "(((((2 ^ 3) * (4 + 5)) - 6) / 7) * (8 + 9))",
			err: false,
		},
		{
			name:        "Invalid expression (not enough operands)",
			input:  	 "3 +",
			expected:    "",
			err: true,
		},
		{
			name:        "Invalid token",
			input:  	 "4 2 x * 5 +",
			expected:    "",
			err: true,
		},
		{
			name:        "Empty expression",
			input:  	 "",
			expected:    "",
			err: true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result, err := PostfixToInfix(tt.input)
			if tt.err {
				assert.Error(t, err, fmt.Sprintf("Test %d: Expected error", i+1))
			} else {
				assert.NoError(t, err, fmt.Sprintf("Test %d: Unexpected error", i+1))
				assert.Equal(t, tt.expected, result, fmt.Sprintf("Test %d: Unexpected result", i+1))
			}
		})
	}
}

func ExamplePostfixToInfix() {
	test := "3 4 + 5 *"

	result, err := PostfixToInfix(test)
	fmt.Printf("%v, %v\n", result, err)
	// Output:
	// ((3 + 4) * 5), <nil>
}
