package leetcode

// https://leetcode.com/problems/insert-interval/description/

// no overlap [1,2],[5,6] ins [3,4] => [1,2][3,4][5,6]
// left overlap [1,5],[7,9] ins [2,6] => [1,6][7,9]
// right overlap [1,2][5,9] ins [3,6] => [1,2][3,9]
// outer overlap [1,3][6,9] ins [2,7] => [1,9]
// inner overlap [2,3][4,5] ins [1,6] => [1,6]

// We can iterate all intervals, merging when necessary, and put the non overlapping ones in the correct order
func insert(intervals [][]int, newInterval []int) [][]int {
	merge := func(a []int, b []int) []int {
		return []int{min(a[0], b[0]), max(a[1], b[1])}
	}

	// using as a stack
	answer := [][]int{newInterval}
	for _, tup := range intervals {
		// pop answer stack
		prev := answer[len(answer)-1]
		answer = answer[:len(answer)-1]
		// no overlap
		if tup[1] < prev[0] {
			answer = append(answer, tup)
			answer = append(answer, prev)
		} else if tup[0] > prev[1] {
			answer = append(answer, prev)
			answer = append(answer, tup)
		} else {
			merged := merge(tup, prev)
			answer = append(answer, merged)
		}

	}
	return answer
}
