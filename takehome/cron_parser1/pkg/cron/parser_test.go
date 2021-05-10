package cron

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		input         string
		expected      *Schedule
		expectedError string
	}{
		{
			name:  "readme example",
			input: "*/15 0 1,15 * 1-5 /usr/bin/find",
			expected: &Schedule{
				Minutes:     []int{0, 15, 30, 45},
				Hours:       []int{0},
				DaysOfMonth: []int{1, 15},
				Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DaysOfWeek:  []int{1, 2, 3, 4, 5},
				Command:     "/usr/bin/find",
			},
			expectedError: "",
		},
		{
			name:  "combining range step glob and one",
			input: "1,30-36/3,15-18,*/20 0 1 1 1 /usr/bin/find",
			expected: &Schedule{
				Minutes:     []int{0, 1, 15, 16, 17, 18, 20, 30, 33, 36, 40},
				Hours:       []int{0},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "/usr/bin/find",
			},
			expectedError: "",
		},
		{
			name:  "repeating numbers",
			input: "1,2,3,1-3/1 0 1 1 1 /usr/bin/find",
			expected: &Schedule{
				Minutes:     []int{1, 2, 3},
				Hours:       []int{0},
				DaysOfMonth: []int{1},
				Months:      []int{1},
				DaysOfWeek:  []int{1},
				Command:     "/usr/bin/find",
			},
			expectedError: "",
		},
		{
			name:          "over limit",
			input:         "0 0 1 13 1 /usr/bin/find",
			expectedError: "expected number in range of 1 to 12, got: 13",
		},
		{
			name:          "range over limit",
			input:         "0 0 1 2-13 1 /usr/bin/find",
			expectedError: "expected maximum of 12, got: 2-13",
		},
		{
			name:          "unexpected glob",
			input:         "0 0 1 *-12 1 /usr/bin/find",
			expectedError: "expected range from numbers, got: *-12",
		},
		{
			name:          "wrong format",
			input:         "123456 cmd",
			expectedError: "Wrong format of crontab",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			parser, _ := NewParser(&ParserOptions{})

			actual, err := parser.Run(tc.input)

			if tc.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "errors not matching")

				return
			}

			assert.Equal(t, tc.expected, actual, "wrong parsing result")
		})
	}
}
