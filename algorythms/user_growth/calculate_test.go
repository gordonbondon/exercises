package user_growth

import (
	"testing"
)

func TestGetBillionUsersDay(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		target   float64
		input    []float64
		expected int
	}{
		{
			name:     "52",
			target:   float64(1000000000),
			input:    []float64{1.5},
			expected: 52,
		},
		{
			name:     "79",
			target:   float64(1000000000),
			input:    []float64{1.1, 1.2, 1.3},
			expected: 79,
		},
		{
			name:     "1047",
			target:   float64(1000000000),
			input:    []float64{1.01, 1.02},
			expected: 1047,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := GetBillionUsersDay(tc.target, tc.input)

			if tc.expected != result {
				t.Errorf("expected: %d, got %d", tc.expected, result)
			}
		})
	}
}
