package rotate_cypher

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    string
		rot      int
		expected string
	}{
		{
			name:     "simple",
			input:    "abcdefghijklmNOPQRSTUVWXYZ0123456789",
			rot:      39,
			expected: "nopqrstuvwxyzABCDEFGHIJKLM9012345678",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Encrypt(tc.input, tc.rot)

			if tc.expected != result {
				t.Errorf("expected: %s, got %s", tc.expected, result)
			}
		})
	}
}
