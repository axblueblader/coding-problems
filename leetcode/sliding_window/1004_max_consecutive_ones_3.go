package leetcode

// https://leetcode.com/problems/max-consecutive-ones-iii/

// Naive
// Iterate the array from beginning
// When i see a zero, i'll flip it immediately to continue forward
// If i have flipped K zeros, then i cannot conitnue
// Compare the current subarray length to max to get final answer
// The move the pointer up by 1, and repeat this process

// However notice this will make me repeat a lot of positions
// Changing the amount of zeroes flipped is enough to know when the subarray ends
// We don't have to recheck the position is 0 or 1, since we've passed through it already

// Have to keep track of which position our subarray should start next
func longestOnes(nums []int, k int) int {
	lastZero := 0
	firstZero := 0
	flipCount := 0
	currentLen := 0
	res := 0
	i := 0
	for i = 0; flipCount < k && i < len(nums); i++ {
		if nums[i] == 1 {
			currentLen++
		} else {
			if flipCount == 0 {
				firstZero = i
			}
			flipCount++
			currentLen++

			nums[lastZero] = i
			lastZero = i
		}
	}
	for ; i < len(nums); i++ {
		if nums[i] == 1 {
			currentLen++
		} else if k != 0 {
			res = max(res, currentLen)
			newFirstZero := nums[firstZero]
			currentLen = i - newFirstZero + (newFirstZero - firstZero)
			firstZero = newFirstZero
			nums[lastZero] = i
			lastZero = i
		} else {
			res = max(res, currentLen)
			currentLen = 0
		}
	}
	return max(res, currentLen)
}

// A better implementation, the idea is the same but instead of keeping track of the previous zero positions
// We use a for loop to go back based on the current subarray size
func longestOnesBetter(nums []int, k int) int {
	used := 0
	size := 0
	best := 0

	for i, x := range nums {
		if x != 0 {
			size++
			if best < size {
				best = size
			}
		} else if used < k {
			used++
			size++
			if best < size {
				best = size
			}
		} else {
			for ; nums[i-size] == 1; size-- {
			}
		}
	}
	return best
}
