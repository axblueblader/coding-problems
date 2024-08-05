package leetcode

// https://leetcode.com/problems/minimum-size-subarray-sum/description/

import "math"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	minLen := math.MaxInt
	start := 0
	for i, val := range nums {
		sum += val

		for ; sum >= target && start <= i; start++ {
			if minLen > i-start+1 {
				minLen = i - start + 1
			}
			sum -= nums[start]
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
