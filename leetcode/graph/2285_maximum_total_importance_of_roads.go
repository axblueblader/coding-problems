package leetcode

import "slices"

// https://leetcode.com/problems/maximum-total-importance-of-roads/description/
//
// A city should have more value if there's more roads connected to
// so optimally we should count the roads a city connect, sort it, and assign value from high to low
// Naive steps:
// count the roads connected to city
// sort city by the count, this will also be the city's number assignment
// calculate the answer based on the assignment by count * assign_number
// e.g
// number of occurence:
// [2 3 4 2 1] meaning node 0 is connected to 2 roads
// sorted
// [4 3 2 2 1]
// then times it with
// 4*5 + 3*4 + 2*3 + 2*2 + 1
// will get final result
func maximumImportance(n int, roads [][]int) int64 {
	connected := make([]int, n)
	for _, pair := range roads {
		connected[pair[0]]++
		connected[pair[1]]++
	}
	slices.Sort(connected)
	sum := int64(0)
	for i, count := range connected {
		sum += int64((i + 1) * count)
	}
	return sum
}
