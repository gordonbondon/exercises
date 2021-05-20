package string_compression1

import (
	"strconv"
)

func CompressString(chars []byte) int {
	groupchar := byte(0)
	groupcount := 0
	writepos := 0

	for i, c := range chars {
		if groupchar != c {
			// when char changes
			groupchar = c
			groupcount = 1
		} else {
			// when we see the same char as before
			groupcount++
		}

		// if next char is different or we are at the end
		if i+1 == len(chars) || chars[i+1] != groupchar {
			chars[writepos] = groupchar
			writepos++

			if groupcount > 1 {
				for i, b := range []byte(strconv.Itoa(groupcount)) {
					chars[writepos+i] = b
				}

				writepos = writepos + len([]byte(strconv.Itoa(groupcount)))
			}
		}
	}

	return writepos
}
