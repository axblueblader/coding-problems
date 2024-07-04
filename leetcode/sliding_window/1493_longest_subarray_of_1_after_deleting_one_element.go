package leetcode

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/

// Best to delete 0s
// [0,1,1,1,0,1,1,0,1]
// [0,x,x,x] - save as best so far
//
//	[x,x,x,0] - start next count from next of previous zero
//	[x,x,x,x,x,x] - save best so far
//	        [x,x,0,x]
func longestSubarray(nums []int) int {
	best := 0
	curr := 0
	lastZeroIndex := -1
	for i, val := range nums {
		if val != 0 {
			curr++
			best = max(best, curr)
		} else {
			if lastZeroIndex != -1 {
				curr = i - lastZeroIndex - 1
			}
			lastZeroIndex = i
		}
	}
	if lastZeroIndex == -1 {
		return best - 1
	}
	return best
}
