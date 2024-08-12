package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

// sell multiple days
// so we just need to keep track of an increasing sequence
func maxProfit(prices []int) int {
	maxProfit := 0
	start := 0
	end := 1
	for end < len(prices) {
		// fmt.Println(start, end, prices[start], prices[end], maxProfit)
		if prices[end] > prices[end-1] {
			end++
		} else {
			maxProfit += prices[end-1] - prices[start]
			start = end
			end++
		}
	}

	if start < end {
		maxProfit += prices[end-1] - prices[start]
	}

	return maxProfit
}
