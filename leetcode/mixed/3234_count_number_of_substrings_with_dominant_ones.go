package leetcode

// https://leetcode.com/problems/count-the-number-of-substrings-with-dominant-ones/description/

import "math"

// brute force:
// for start to len(s)
//  for end=start to len(s)
//   for i=start to end check by counting s[i]

//	1 0 1 1 0 1
//
// 0 1 1 2 3 3 4
//
//	x
//
// s   e
// s         e
//
//	s   e

func numberOfSubstrings(s string) int {
	sum := 0
	prefix := make([]int, len(s)+1)
	for i, ch := range s {
		if ch == '1' {
			sum += 1
		}
		prefix[i+1] = sum
	}
	// fmt.Println(prefix)

	count := 0
	for left := 0; left < len(s); left++ {
		for right := left; right < len(s); right++ {
			ones := prefix[right+1] - prefix[left]
			zeros := right - left + 1 - ones
			if zeros*zeros == ones {
				count++
			} else if zeros*zeros < ones {
				count++

				// i can take this much more zeros
				diff := int(math.Sqrt(float64(ones))) - zeros
				nextRight := right + diff

				if nextRight >= len(s) {
					// add all combination till the end
					count += len(s) - right - 1
				} else {
					count += diff
				}
				right = nextRight
			} else {
				// i need at least this much more 'ones' to match zeros*zeros
				right += zeros*zeros - ones - 1
			}
			// fmt.Println(left, right, ones, zeros, count)
		}
	}

	return count
}
