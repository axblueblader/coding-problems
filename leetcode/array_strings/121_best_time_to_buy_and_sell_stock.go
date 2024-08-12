package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
//
import "math"

// can only buy and sell once
// max profit will be next days - current min day so far
func maxProfit(prices []int) int {
	minDay := math.MaxInt
	maxProfit := 0

	for _, val := range prices {
		if val > minDay {
			maxProfit = max(maxProfit, val-minDay)
		} else {
			minDay = val
		}
	}

	return maxProfit
}
