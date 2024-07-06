package leetcode

// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

// I have "abc", then an anagram needs to contain "a", "b", "c"
// Store the character needed in a map
// a substring of 's' is an anagram if all characters inside, is also inside 'p'
// exists[s[i]] == true
// where does the substring ends?
// when it has contains all the unique characters of 'p'
// this condition translate to if frequency map of p matches frequency map of subarray
// then the subarray is an anagram
// However recomputing the frequency map will result in time limit
// we must use sliding window method, decreasing the frequency whenever start index needs to move
func findAnagrams(s string, p string) []int {
	appear := make(map[byte]int)
	for _, ch := range p {
		appear[byte(ch)]++
	}
	res := []int{}
	seen := make(map[byte]int)
	startIndex := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if appear[ch] == 0 {
			startIndex = i + 1
			clear(seen)
			continue
		}
		if seen[ch]+1 > appear[ch] {
			j := startIndex
			for ; s[j] != s[i]; j++ {
				seen[s[j]]--
			}
			startIndex = j + 1
			continue
		}
		seen[ch]++
		if len(seen) != len(appear) {
			continue
		}
		isAnagram := true
		for k, val := range appear {
			if seen[k] != val {
				isAnagram = false
				break
			}
		}
		if isAnagram {
			res = append(res, startIndex)
			seen[s[startIndex]]--
			startIndex++
		}
	}
	return res
}
