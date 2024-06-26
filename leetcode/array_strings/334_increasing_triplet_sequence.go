package leetcode

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence/description/
//
// A good way to solve this is start from naive 3 nested loop
//
//	for i {
//	 for j=i+1 {
//	  for k=j+1 {
//	    if nums[i] < nums[j] < nums[j] return true
//	  }
//	 }
//	}
//
// Then try to see what can be cache instead of looping through
// If we iterate from the left, we can keep track of the smallest number so far
// That would reduce it to 2 nested loop
//
// > minLeft:= math.MaxInt32
//
//	for i {
//	 minLeft = min(minLeft, nums[i])
//	 for j=i+1 {
//	  if minLeft < nums[i] < nums[j] return true
//	 }
//	}
//
// The final crux is realizing we can apply the same trick again, but for the 2nd smallest number
// which will reduce solution to 1 loop
func increasingTriplet(nums []int) bool {
	secondSmallest := math.MaxInt32
	smallest := math.MaxInt32
	for _, val := range nums {
		if val <= smallest {
			smallest = val
		} else if val <= secondSmallest {
			secondSmallest = val
		} else {
			return true
		}
	}
	return false
}
