package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "Valid expression",
			input:       "5 1 2 + 4 * + 3 -",
			expected:    "((5 + ((1 + 2) * 4)) - 3)",
			expectedErr: nil,
		},
		{
			name:        "Invalid expression",
			input:       "5 1 + +",
			expected:    "",
			expectedErr: errors.New("invalid expression: not enough operands"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}

			handler := &ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()
			if err != nil && tt.expectedErr == nil {
				t.Fatalf("Неочікувана помилка: %v", err)
			}
			if err == nil && tt.expectedErr != nil {
				t.Fatalf("Очікувалась помилка, але отримано nil")
			}
			if err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error() {
				t.Fatalf("Очікувалась помилка %v, отримано %v", tt.expectedErr, err)
			}

			if tt.expected != "" && output.String() != tt.expected+"\n" {
				t.Fatalf("Очікувався результат %s, отримано %s", tt.expected, output.String())
			}
		})
	}
}
