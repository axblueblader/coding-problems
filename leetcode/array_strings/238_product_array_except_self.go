package leetcode

// https://leetcode.com/problems/product-of-array-except-self/description/
//
// Naive approach: for each position, traverse the whole array to get the product exlucding it
//
// Use caching: traverse the array, storing the product from one direction, in this case, the suffix product, stored from end to start
// Example:
// original: [1,2,3,4]
// suffix: [24,24,12,4]
// Now we can traverse it again from beginning to end, calculatin the prefix product
// At position i-th, the result will be prefix_product * suffix[i+1]
func productExceptSelf(nums []int) []int {
	n := len(nums)
	suffix := make([]int, n)
	for i, prd := n-1, 1; i >= 0; i-- {
		suffix[i] = prd * nums[i]
		prd = prd * nums[i]
	}
	prd := 1
	for i := 0; i < n-1; i++ {
		suffix[i] = prd * suffix[i+1]
		prd = prd * nums[i]
	}
	suffix[n-1] = prd
	return suffix
}
