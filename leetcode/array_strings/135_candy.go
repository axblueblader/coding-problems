package leetcode

// https://leetcode.com/problems/candy/description/

func candy(ratings []int) int {
	dist := make([]int, len(ratings))
	rateNums := make(map[int]bool)
	// every needs one
	for i := range ratings {
		rateNums[ratings[i]] = true
		dist[i]++
	}
	answer := 0
	// Repeat the steps of ditributing more candies to position with higher ratings
	// We know the maximum time we need to do this is based on the number of different ratings
	for range rateNums {
		// A faster way would be to store only the positions that need increasing
		// for [[2,2,0,1,1,3,3]
		// e.g map[1] = [3,4], map[2] = [0,1], map[3] = [5,6]
		// Then we can iterate based on the sorted rating type
		// e.g the types of ratings are [1,4,10,100,...]
		// loop 1: check all position of 1s, increase if needed
		// loop 2: check all position of 4s, increase if needed
		// loop 3: check all position of 10s, increase if needed
		// repeat same up to the highest rating
		for i := 1; i < len(dist); i++ {
			if ratings[i] > ratings[i-1] && dist[i] <= dist[i-1] {
				dist[i]++
			} else if ratings[i] < ratings[i-1] && dist[i] >= dist[i-1] {
				dist[i-1]++
			}
		}
	}
	for _, val := range dist {
		answer += val
	}
	return answer
}
