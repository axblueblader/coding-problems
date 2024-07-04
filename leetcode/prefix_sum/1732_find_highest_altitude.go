package leetcode

// https://leetcode.com/problems/find-the-highest-altitude/description/

// -5 means goes down 5 so current altitude will be current - 5
// +3 means goes up 3, current = current + 3
// keep track of the highest along the way
func largestAltitude(gain []int) int {
	best := 0
	curr := 0
	for _, val := range gain {
		curr = curr + val
		best = max(best, curr)
	}
	return best
}
