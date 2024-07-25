package leetcode

// https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

func maxScore(cardPoints []int, k int) int {
	arrSize := len(cardPoints) - k

	sum := 0
	total := 0
	for i := 0; i < arrSize; i++ {
		sum += cardPoints[i]
		total += cardPoints[i]
	}
	minSumArr := sum

	for i := arrSize; i < len(cardPoints); i++ {
		total += cardPoints[i]
		sum = sum + cardPoints[i] - cardPoints[i-arrSize]
		minSumArr = min(minSumArr, sum)
	}

	return total - minSumArr
}
