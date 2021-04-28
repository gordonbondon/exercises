package check

import (
	"strings"
)

// Check if you can make the result string by copy-pasting the whole input string
// multiple times and inserting it in any place
func Check(input string, result string) bool {
	solution := check(input, result)

	return solution
}

func check(input string, result string) bool {
	scratch := result
	solution := false

	for {
		split := strings.Split(scratch, input)

		if len(split) == 1 && split[0] == scratch {
			// string not found in result string, so we can not continue
			break
		}

		scratch = strings.Join(split, "")

		if scratch == "" || scratch == input {
			solution = true
			break
		}
	}

	return solution
}
