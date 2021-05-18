package valid_parentheses

import "strings"

var (
	openToClose = map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	closeToOpen = map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
)

func ValidateParentheses(input string) bool {
	brackets := strings.Split(input, "")

	open := make([]string, 0)

	for _, br := range brackets {
		if _, ok := openToClose[br]; ok {
			open = append(open, br)

			continue
		}

		if op, ok := closeToOpen[br]; ok {
			if len(open) == 0 || open[len(open)-1] != op {
				return false
			}

			open = open[:len(open)-1]

			continue
		}

		// not a bracket
	}

	if len(open) > 0 {
		return false
	}

	return true
}
