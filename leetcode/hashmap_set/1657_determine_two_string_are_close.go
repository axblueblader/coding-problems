package leetcode

// https://leetcode.com/problems/determine-if-two-strings-are-close/

// different length then 'false'
// meaning of operation 1: contains the same set of characters
// operation 2: contains the same pattern of occurence
// word1 = "aaeeabcbd" -> 3a 2b 2e 1d 1c
// word2 = "ebacbdcec" -> 3c 2b 2e 1d 1a
// occurrences is 2-1s, 2-2s, 1-3s
// => same number of occurrences
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	occur1 := make(map[rune]int)
	occur2 := make(map[rune]int)
	for _, ch := range word1 {
		occur1[ch]++
	}
	for _, ch := range word2 {
		occur2[ch]++
	}
	// same set of characters
	for k := range occur1 {
		if occur2[k] == 0 {
			return false
		}
	}
	count1 := make([]int, len(word1)+1)
	count2 := make([]int, len(word2)+1)
	for _, occ := range occur1 {
		count1[occ]++
	}
	for _, occ := range occur2 {
		count2[occ]++
	}
	// same count of number of occurence
	for i, v := range count1 {
		if v != count2[i] {
			return false
		}
	}
	return true
}
