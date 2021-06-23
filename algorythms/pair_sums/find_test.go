package pair_sums

import (
	"testing"
)

func TestFindNumberOfWays(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		target   int
		input    []int
		expected int
	}{
		{
			name:     "duplicate",
			target:   6,
			input:    []int{1, 5, 3, 3, 3},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := FindNumberOfWays(tc.input, tc.target)

			if tc.expected != result {
				t.Errorf("expected: %d, got %d", tc.expected, result)
			}
		})
	}
}
