package leetcode

// https://leetcode.com/problems/unique-number-of-occurrences/description/

// Count the occurrences
// then use a different map to check for their uniqueness
func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	for _, val := range arr {
		occurrences[val]++
	}
	seen := make(map[int]bool)
	for _, occ := range occurrences {
		if seen[occ] {
			return false
		} else {
			seen[occ] = true
		}
	}
	return true
}
