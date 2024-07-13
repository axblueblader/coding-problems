package leetcode

// https://leetcode.com/problems/decode-string/

import (
	"strconv"
	"strings"
)

// base case 3[abc], see 3, remember it, store string until ] is met -> abc, duplicate it 3 times
// multi digits number case 100[abc] -> missed this case at the beginning
// nested case 3[a4[b]], the string inside need to be decoded and combined with previous string
// doesn't allow number inside [] so we have less cases to check
func decodeString(s string) string {
	stack := []string{}
	// buffer for number
	digits := ""
	for _, ch := range s {
		if '0' <= ch && ch <= '9' {
			digits = digits + string(ch)
			continue
		}
		if ch == '[' {
			stack = append(stack, digits)
			stack = append(stack, string(ch))
			digits = ""
			continue
		}
		if ch != ']' {
			stack = append(stack, string(ch))
			continue
		}
		// this is ]
		str := ""
		for len(stack) > 0 && stack[len(stack)-1] != "[" {
			// need to plus in reversed order
			str = stack[len(stack)-1] + str
			stack = stack[:len(stack)-1]
		}
		// pop the [ and number
		// fmt.Println(stack, str)
		if len(stack) > 0 {
			num, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			decoded := ""
			for range num {
				decoded += str
			}
			stack = append(stack, decoded)
		}
	}
	return strings.Join(stack, "")
}
