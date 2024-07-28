package leetcode

// https://leetcode.com/problems/top-k-frequent-elements/

// most frequent elements
// calculate frequency
// map[num]frequency

// invert map to map[frequency][]nums
// bucket sort instead of using map
// 1 1 2 2 3
// arr[1] -> 3
// arr[2] -> 1,2
// loop through this arr, reducing k

func topKFrequent(nums []int, k int) []int {
	numToFreq := make(map[int]int)

	for _, val := range nums {
		numToFreq[val]++
	}

	// fmt.Println(numToFreq)

	freqToNums := make([][]int, len(nums)+1)

	for key, val := range numToFreq {
		if freqToNums[val] == nil {
			freqToNums[val] = make([]int, 0)
		}
		freqToNums[val] = append(freqToNums[val], key)
	}

	// fmt.Println(freqToNums)
	ans := []int{}
	for i := len(freqToNums) - 1; i >= 0 && k > 0; i-- {
		if len(freqToNums[i]) > 0 {
			for _, val := range freqToNums[i] {
				ans = append(ans, val)
				k--
				// fmt.Println(ans, k)
				if k == 0 {
					break
				}
			}
		}
	}

	return ans
}
