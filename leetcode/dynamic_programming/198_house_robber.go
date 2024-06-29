package leetcode

// https://leetcode.com/problems/house-robber/

// Naive approach using recursive
//
// At every position, we need to choose whether to pick the current house or not
// Picking a house means we cannot pick the next one, and can only pick the +2 one
// To choose we must calculate all the cases until the end, planning forward
// This can be written as a recursive equation as
// f(i) = max(nums[i] + f(i+2), f(i+1))
// or f(i) max(nums[i] + f(i-2), f(i-1)) if we start from the end
func robRecursion(nums []int) int {
	n := len(nums)
	var robAt func(m int) int
	robAt = func(i int) int {
		if i >= n {
			return 0
		}
		return max(nums[i]+robAt(i+2), robAt(i+1))
	}
	return robAt(0)
}

// Memoization aka Caching
//
// Notice there is a repetitive step, for example
// at i=0, we calculate i+1=1 and i+2=2
// at i=1, we calculate i+1=2 and i+2=3
// The value at i=2 is calculated twice
// This means if we store it, we can avoid the computation
func robRecursionCached(nums []int) int {
	n := len(nums)
	cache := make([]int, n+2)
	var robAt func(i int) int
	robAt = func(i int) int {
		if i >= n {
			return 0
		}
		if cache[i] != 0 {
			return cache[i]
		}
		cache[i+2] = robAt(i + 2)
		cache[i+1] = robAt(i + 1)
		cache[i] = max(nums[i]+cache[i+2], cache[i+1])
		return cache[i]
	}
	return robAt(0)
}

// Tabulation
//
// If we think about it, recursion is another way of repeting steps, another way of doing loops
// In the previous method, we start at i=0,
// but due to recursion, we actually calculate values from the END first
// cache[0] is actually the LAST value to be calculated
// This is similiar to doing a for loop from END back to 0
//
//	for i:=n-1;i>=0;i-- {
//	  cache[i] = max(nums[i] + cache[i+2],cache[i+1])
//	}
func robDP(nums []int) int {
	n := len(nums)
	bestSoFar := make([]int, n+2)
	for i := n - 1; i >= 0; i-- {
		bestSoFar[i] = max(bestSoFar[i+2]+nums[i], bestSoFar[i+1])
	}
	return bestSoFar[0]
}

// Optimized memory
//
// Yet there is still another improvement can be made
// In the loop, notice the furthest we need to store is 2 position away
// Therefore we can store them in a variable instead of needing an array
func robDPWithSimpleVariables(nums []int) int {
	var bestHere, bestAtPlusOne, bestAtPlusTwo int
	for i := len(nums) - 1; i >= 0; i-- {
		bestHere = max(bestAtPlusTwo+nums[i], bestAtPlusOne)
		tmp := bestAtPlusOne
		bestAtPlusOne = bestHere
		bestAtPlusTwo = tmp
	}
	return max(bestHere, bestAtPlusOne)
}
