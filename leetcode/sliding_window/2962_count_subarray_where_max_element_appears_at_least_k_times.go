package leetcode

// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

func countSubarrays(nums []int, k int) int64 {
	curMax := 0
	freq := 0
	count := 0

	for _, val := range nums {
		if curMax < val {
			curMax = val
		}
	}

	left := 0
	right := 0
	for left < len(nums) {
		if right < len(nums) && nums[right] == curMax {
			freq++
		}
		if right < len(nums) {
			right++
		} else {
			left++
			freq--
		}
		if freq < k {
			if right >= len(nums) {
				break
			}
			continue
		}
		if freq > k {
			freq--
			left++
		}
		for left < len(nums) && nums[left] != curMax {
			left++
		}
		// fmt.Println(left, right, nums[left:right], freq)
		count++
		count += left

	}

	return int64(count)
}
