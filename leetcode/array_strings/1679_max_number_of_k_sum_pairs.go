package leetcode

// https://leetcode.com/problems/max-number-of-k-sum-pairs/description


// Naive
// Try for all if it's equal to K, then remove
// for i to n
//
//	 if nums[i] == 0 skip
//	for j=i+1 to n
//	 if nums[j] == 0 skip
//	 if nums[i] + nums[j] == k
//	  count++
//	  nums[i] = 0
//	  nums[j] = 0

// Starting from left, when we are at position i-th
// We know exactly which number we need already
// e.g k = 4, num[0] = 3
// we know if there's a num[i] = 1, we can count 1 up
// hence a map can be used by storing map[k-nums[i]] when we visit i
// in the case of
// k = 4
// nums = [4,4,1,1]
// Notice we also need to increase map[k-nums[i]] for every match pair 
// or else it will miss out pairs if using 0 or boolean
// Time O(N) Space O(N)
func maxOperationsWithMap(nums []int, k int) int {
	pairAt := make(map[int]int)
	count := 0
	for _, val := range nums {
		if pairAt[val] != 0 {
			count++
			pairAt[val] -= 1
			continue
		}
		pairAt[k-val] += 1
	}
	return count
}

// Alternatively there's a 2 pointer solution
// Look from a different perspective, at position i-th, what we're trying to do
// is search for the next number so that nums[i] + nums[j] = k
// When searching numbers, it's faster to do if the array is sorted
// After sorting we can use 2 opposite pointers to find a pair
// By moving the smaller pointer, increasing the sum, if sum is smaller than k
// or move the bigger pointer, lessening the sum, if it's larger than k
// Time O(NlogN) due to sort, Space O(1)
func maxOperationsTwoPointers(nums []int, k int) int {
	left,right:=0,len(nums)-1
	count := 0
    slices.Sort(nums)
    for left < right {
        if nums[left] + nums[right] > k {
            right--
        } else if nums[left] + nums[right] < k {
            left++
        } else {
            count++
            left++
            right--
        }
    }
	return count
}
