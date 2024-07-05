package leetcode

// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/

// Basically implement Set and 'set difference' operation
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for _, val := range nums1 {
		set1[val] = true
	}
	for _, val := range nums2 {
		set2[val] = true
	}
	answer1 := []int{}
	answer2 := []int{}
	for k := range set1 {
		if !set2[k] {
			answer1 = append(answer1, k)
		}
	}
	for k := range set2 {
		if !set1[k] {
			answer2 = append(answer2, k)
		}
	}

	return [][]int{answer1, answer2}
}
