package leetcode

import (
	"math/rand"
	"slices"
	"sort"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/

// There are multiple ways to solve this: quickselect, heap
// But here i'm using an easy to follow map + sort approach that sorts the key instead of the whole array
func findKthLargest(nums []int, k int) int {
	occur := make(map[int]int)
	for _, val := range nums {
		occur[val]++
	}
	keys := []int{}
	for key := range occur {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	slices.Reverse(keys)
	// values that occurred, sorted in descending order
	for _, key := range keys {
		if occur[key] > k {
			return key
		}
		k = k - occur[key]
		if k == 0 {
			return key
		}
	}
	// actually never reach
	return 0
}

// Quick select is similiar to quick sort
// Pick a random index to divide the array into 2
// Sort so that all elements before the index is smaller
// all elements after the index is larger
// If the index == k, we found the answer
// If not, pick the next subarray based on k is smaller or larger and repeat pivot+partition step
func findKthLargestQuickSelect(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		pivotIndex := left + rand.Intn(right-left+1)
		newPivotIndex := partition(nums, left, right, pivotIndex)
		if newPivotIndex == len(nums)-k {
			return nums[newPivotIndex]
		} else if newPivotIndex > len(nums)-k {
			right = newPivotIndex - 1
		} else {
			left = newPivotIndex + 1
		}
	}
}

func partition(nums []int, left int, right int, pivotIndex int) int {
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storedIndex := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storedIndex] = nums[storedIndex], nums[i]
			storedIndex++
		}
	}
	nums[right], nums[storedIndex] = nums[storedIndex], nums[right]
	return storedIndex
}
