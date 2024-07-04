package leetcode

// https://leetcode.com/problems/find-pivot-index/

// Naive:
// Calculate all the left sum
// Calculate all the right sum
// find the index 'i' where leftSum[i-1] == rightSum[i+1]

// One way to reduce memory usage is to calculate the rightSum starting and store it in 1 variable
// During iteration from the left we can get the rightSum by subtracting the current value from it
func pivotIndex(nums []int) int {
	left := make([]int, len(nums)+2)
	right := make([]int, len(nums)+2)

	for i, j := 0, len(nums)-1; i < len(nums)-1; i, j = i+1, j-1 {
		left[i+1] = left[i] + nums[i]
		right[j] = right[j+1] + nums[j]
	}

	for i := 0; i < len(nums); i++ {
		if left[i] == right[i+1] {
			return i
		}
	}
	return -1
}
