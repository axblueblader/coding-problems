package leetcode

// https://leetcode.com/problems/subarray-product-less-than-k/

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	n := len(nums)
	count := 0
	left := 0
	right := 0
	prd := 1
	for right < n {
		prd = prd * nums[right]
		for prd >= k {
			prd = prd / nums[left]
			left++
		}
		count += right - left + 1
		right++
	}

	return count
}
