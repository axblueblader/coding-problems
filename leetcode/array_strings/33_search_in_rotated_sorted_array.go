package leetcode

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

// [0,1,2,4,5,6,7]
// [7,0,1,2,4,5,6]
// [2,4,5,6,7,0,1]
// [5,6,7,0,1,2,4]
// left most is larger than right most => there's a pivot inside this region, search both sides
func search(nums []int, target int) int {
	var binarySearch func(int, int) int
	binarySearch = func(left int, right int) int {
		if left > right {
			return -1
		}
		mid := left + (right-left)/2
		// fmt.Println(right, left, mid, nums[left:right+1], nums[mid])
		if nums[mid] == target {
			return mid
		}
		if nums[left] > nums[right] {
			return max(binarySearch(left, mid-1), binarySearch(mid+1, right))
		}
		if target > nums[mid] {
			return binarySearch(mid+1, right)
		} else {
			return binarySearch(left, mid-1)
		}
	}
	return binarySearch(0, len(nums)-1)
}
