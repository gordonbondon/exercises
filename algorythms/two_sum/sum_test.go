package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "test one",
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "test two",
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "test three",
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			t.Logf("array: %v, target: %d", tc.input, tc.target)

			result := TwoSum(tc.input, tc.target)

			assert.Equal(t, tc.expected, result, "wrong output array")
		})
	}
}
