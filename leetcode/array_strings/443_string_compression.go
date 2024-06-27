package leetcode

import "strconv"

// https://leetcode.com/problems/string-compression/description/
//
// We need to store the index of where we can write next
// But the big GOTCHA is since it's a byte array, when we store the count, it must be convert to a string
// and the write index must increase accordingly
// e.g count = 10 means we need 2 space to write
// count = 100 means we need 2 space
// Need an 'if' check if the count is only 1, since we don't write the count when it's 1
// Need more 'if' checks for the final position

func compress(chars []byte) int {
	nextChPos := 0
	curChar := chars[0]
	count := 0
	for i := 0; i <= len(chars); i++ {
		if i == len(chars) || curChar != chars[i] {
			chars[nextChPos] = curChar
			nextChPos++
			if count > 1 {
				countStr := strconv.Itoa(count)
				for _, char := range countStr {
					chars[nextChPos] = byte(char)
					nextChPos++
				}
				count = 1
			}
			if i < len(chars) {
				curChar = chars[i]
			}
		} else {
			count++
		}
	}

	return nextChPos
}
