package leetcode

// https://leetcode.com/problems/container-with-most-water/

// Maximum amount water = max area, lower height * distance
//
// Naive:
// for each start:
//  for each end:
//   result = max(result, min(height[start],height[end]) * (end - start))
// O(n^2)

// Because higher distance is better, let's begin at 0 and END
// To get the maximum area, the difference between the height should be smallest
// To get the maximum area, the height should be highest
// So after each calculation, we should move the smaller ones to see if we can find a better area
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}
