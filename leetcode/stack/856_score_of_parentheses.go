package leetcode

// https://leetcode.com/problems/score-of-parentheses/

// ()
// 1
// ((()()())())
// O O 1 1 1 C
// O 6 1 C
// 14
func scoreOfParentheses(s string) int {
	stack := []int{}

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, 0)
		}
		if ch == ')' {
			sum := 0
			for len(stack) > 0 && stack[len(stack)-1] != 0 {
				pop := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				sum += pop
			}
			// if sum == 0 (no inner parans) then plus 1
			sum = max(sum*2, 1)
			// pop the 0
			stack = stack[0 : len(stack)-1]
			// add the sum
			stack = append(stack, sum)
		}
	}
	ans := 0
	for _, val := range stack {
		ans += val
	}
	return ans
}
