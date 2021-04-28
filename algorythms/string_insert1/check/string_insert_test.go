package check

import (
	"testing"
)

func TestCheck(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input            string
		result           string
		expectedSolution bool
	}{
		{
			input:            "XY",
			result:           "XXXYYY",
			expectedSolution: true,
		},
		{
			input:            "XY",
			result:           "XXXYYX",
			expectedSolution: false,
		},
		{
			input:            "XY",
			result:           "XXYXYXYY",
			expectedSolution: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.input+"_"+tc.result, func(t *testing.T) {
			t.Parallel()

			solution := Check(tc.input, tc.result)

			if solution != tc.expectedSolution {
				t.Fail()
			}
		})
	}
}
