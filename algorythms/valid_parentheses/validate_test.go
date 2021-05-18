package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateParentheses(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "good",
			input:    "()",
			expected: true,
		},
		{
			name:     "longer good",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "bad",
			input:    "(]",
			expected: false,
		},
		{
			name:     "longer bad",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "nested good",
			input:    "{[]}",
			expected: true,
		},
		{
			name:     "good with other text",
			input:    "{[testing]}",
			expected: true,
		},
		{
			name:     "missing closure",
			input:    "{[]",
			expected: false,
		},
		{
			name:     "missing open",
			input:    "()]",
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			t.Logf("input: %s", tc.input)

			result := ValidateParentheses(tc.input)

			assert.Equal(t, tc.expected, result, "wrong result")
		})
	}
}
