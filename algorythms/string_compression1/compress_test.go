package string_compression1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateParentheses(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          []byte
		expected       []byte
		expectedLength int
	}{
		{
			name:           "three groups",
			input:          []byte("aabbccc"),
			expected:       []byte("a2b2c3"),
			expectedLength: 6,
		},
		{
			name:           "one char",
			input:          []byte("a"),
			expected:       []byte("a"),
			expectedLength: 1,
		},
		{
			name:           "one char and group",
			input:          []byte("abbbbbbbbbbbb"),
			expected:       []byte("ab12"),
			expectedLength: 4,
		},
		{
			name:           "groups of similar chars",
			input:          []byte("aabbaa"),
			expected:       []byte("a2b2a2"),
			expectedLength: 6,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := CompressString(tc.input)

			assert.Equal(t, tc.expected, tc.input[:result], "wrong result")
			assert.Equal(t, tc.expectedLength, result, "wrong length")
		})
	}
}
