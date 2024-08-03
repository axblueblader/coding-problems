package leetcode

// https://leetcode.com/problems/find-if-digit-game-can-be-won/

func canAliceWin(nums []int) bool {
	singleSum := 0
	doubleSum := 0
	for _, val := range nums {
		if val >= 10 {
			doubleSum += val
		} else {
			singleSum += val
		}
	}

	return doubleSum != singleSum
}
