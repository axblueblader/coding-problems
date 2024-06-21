package leetcode

// Naive way would be to iteraate through each, give extra candies to kid at i-th then traverse array to check for max
// Important observation is the condition that adding the candies will always be max if it's larger or equal the original max among all kids
func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	result := make([]bool, n)
	max := 0
	for i := 0; i < n; i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	for i := 0; i < n; i++ {
		if extraCandies+candies[i] >= max {
			result[i] = true
		}
	}

	return result
}
