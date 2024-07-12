package leetcode

// https://leetcode.com/problems/removing-stars-from-a-string/description/

// the word retains order, so first see, first in
// when we see '*', we need to remove the 'last element'
// this is pattern of a 'stack' First In Last Out
func removeStars(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if ch != '*' {
			stack = append(stack, ch)
		} else {
			// remove last element
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(stack)
}
