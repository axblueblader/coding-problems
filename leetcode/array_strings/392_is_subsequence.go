package leetcode

// https://leetcode.com/problems/is-subsequence/

// Key point is the character has to appear in the right order
// for
// "abc"
// "xxxbxxxcxxaxxxbxc"
// when we find "a" in string 't' at position 'i', the next position we need to search is 'i+1'
func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}
	sIndex := 0
	for _, tChar := range t {
		if sIndex >= len(s) {
			return true
		}
		if rune(s[sIndex]) == tChar {
			sIndex++
		}
	}
	return sIndex == len(s)
}

// Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109,
// and you want to check one by one to see if t has its subsequence.
// In this scenario, how would you change your code?
// In this follow up, since there will be high amounts of 's', and 's' is more often smaller than 't'
// It will be beneficial to 'cache' what we know of 't'
// Here i'm caching the index of all the letters
//
// Another way to reduce the for loop when checking the index is use a 2D array of letters to position
// x\i :0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
// a | 1 | 1 | 1 | 1 | 2 | 2 | 2 | 2 |
// b | 0 | 1 | 1 | 1 | 1 | 2 | 2 | 2 |
// c | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 1 |
// d | 0 | 0 | 1 | 1 | 1 | 1 | 1 | 2 |
// e | 0 | 0 | 0 | 1 | 1 | 1 | 1 | 1 |
// ...
// z | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
func isSubsequenceMulti(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	// this can be a global
	tJump := make(map[rune][]int)
	for tIndex, tChar := range t {
		tJump[tChar] = append(tJump[tChar], tIndex)
	}

	// this part is what should be called multiple times
	isSubOfT := func(s string) bool {
		prevIndex := -1
		for _, sChar := range s {
			if len(tJump[sChar]) == 0 {
				return false
			}
			tmp := prevIndex
			for i, nextIndex := range tJump[sChar] {
				if nextIndex > prevIndex {
					prevIndex = tJump[sChar][i]
					break
				}
			}
			if tmp == prevIndex {
				return false
			}
		}
		return true
	}
	return isSubOfT(s)
}
