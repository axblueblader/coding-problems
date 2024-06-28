package leetcode

// https://leetcode.com/problems/move-zeroes/description/
//
// Naive implementation:
// Start from beginneing, if we meet a zero, iterate from there to find the first non zero and swap
func moveZeroesNaive(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					tmp := nums[j]
					nums[j] = nums[i]
					nums[i] = tmp
					break
				}
			}
		}
	}
}

// Cleaner way:
// We can lessen a loop by remember where the next non zero should be placed
// Then only need 1 loop to iterate swap non zero into their place
// If the current position is a non zero, it will simply be swap into a latter position
// which maintains original ordering
func moveZeroesClean(nums []int) {
	nonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp := nums[nonZero]
			nums[nonZero] = nums[i]
			nums[i] = tmp
			nonZero++
		}
	}
}
