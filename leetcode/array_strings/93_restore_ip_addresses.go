package leetcode

// https://leetcode.com/problems/restore-ip-addresses/

import (
	"strconv"
	"strings"
)

// 1 0 1 0 2 3

/// x       x
// 1.0.1 0 2.3
// 1.0 1.0 2.3
// 1.0 1 0.2.3 => eliminate

//  x     x
// 1.0 1 0.2 3
//  x   x
// 1.0 1.0 2 3

//    x     x
// 1 0.1 0 2.3
// 1 0.1.0 2.3
// 1 0.1 0.2.3

//	x   x
//
// 1 0.1 0.2 3
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}

	invalidZero := func(section string) bool {
		return section[0] == '0' && len(section) > 1
	}

	invalidRange := func(section string) bool {
		val, _ := strconv.Atoi(section)
		return val > 255
	}

	ans := []string{}
	for left := 1; left <= 3; left++ {
		for right := len(s) - 1; right >= len(s)-3; right-- {
			if right-left < 2 || right-left > 6 {
				continue
			}
			first := s[0:left]
			end := s[right:]
			if invalidZero(first) || invalidZero(end) || invalidRange(first) || invalidRange(end) {
				continue
			}
			for i := left + 1; i < right && i <= left+3; i++ {
				if right-i > 3 {
					continue
				}
				mid1 := s[left:i]
				mid2 := s[i:right]
				if invalidZero(mid1) || invalidZero(mid2) || invalidRange(mid1) || invalidRange(mid2) {
					continue
				}
				// fmt.Println(first, mid1, mid2, end)
				ans = append(ans, strings.Join([]string{first, mid1, mid2, end}, "."))
			}
		}
	}

	return ans
}
