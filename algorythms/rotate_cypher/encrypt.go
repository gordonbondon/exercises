package rotate_cypher

import (
	"strings"
)

func Encrypt(input string, rotationFactor int) string {
	var result strings.Builder

	for _, r := range input {
		switch pos := int(r); {
		case pos >= 65 && pos <= 90:
			// runes A-Z 65-90
			result.WriteRune(rotate(pos, 65, 90, rotationFactor))
		case pos >= 97 && pos <= 122:
			// runes a-z 97-122
			result.WriteRune(rotate(pos, 97, 122, rotationFactor))
		case pos >= 48 && pos <= 57:
			// runes 0-9 48-57
			result.WriteRune(rotate(pos, 48, 57, rotationFactor))
		default:
			result.WriteRune(r)
		}
	}

	return result.String()
}

func rotate(pos, min, max, factor int) rune {
	span := max - min + 1
	oldL := pos - min

	rotL := oldL + factor

	newL := rotL % span

	return rune(min + newL)
}
