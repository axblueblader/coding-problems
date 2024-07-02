package leetcode

// https://leetcode.com/problems/maximum-average-subarray-i/

// Naive:
// We can brute force and check all sub arrays, storing the max
// for i to N
//   for j=i to k-1
//     avg = avg + (nums[j]/k)
//   result = max(result, avg)
// Time O(N*k)

// Notice we're computing everything in the middle part multiple times
// Instead we only needed to remove the 1st element and add the last element in each iteration
// Time O(N)
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	res := 0.0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res = float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k]
		sum = sum + nums[i]
		res = max(res, float64(sum)/float64(k))
	}
	return res
}
