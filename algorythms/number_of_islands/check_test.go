package number_of_islands

import (
	"testing"
)

func TestFindIslands(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    [][]string
		expected int
	}{
		{
			name: "one island",
			input: [][]string{
				{"1", "1", "1", "1", "0"},
				{"1", "1", "0", "1", "0"},
				{"1", "1", "0", "0", "0"},
				{"0", "0", "0", "0", "1"},
			},
			expected: 2,
		},
		{
			name: "three islands",
			input: [][]string{
				{"0", "1", "0", "0", "1"},
				{"1", "1", "0", "0", "0"},
				{"1", "0", "1", "0", "0"},
				{"0", "0", "0", "1", "1"},
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := FindIslands(tc.input)

			if tc.expected != result {
				t.Errorf("expected: %d, got %d", tc.expected, result)
			}
		})
	}
}
