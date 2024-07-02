package leetcode

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

// Iterate the k subarray through s
// For each iteration, we can try to count how many vowels and keep track of max count
// That would repeat a lot of positions
// The key to this problem is keeping track of the vowels,
// how do we know if we have just moved from a vowel to decrease it's count
// Use another array to mark indexes with a vowel

// Actually don't need an array to store, just check with the vowel function again on s[i-k]
// The map will be useful if the vowel check is an expensive operation
func maxVowels(s string, k int) int {
	count := 0
	vowels := make([]bool, len(s))
	for i, ch := range s {
		if isVowel(ch) {
			vowels[i] = true
		}
	}
	for i := 0; i < k; i++ {
		if vowels[i] {
			count++
		}
	}

	res := count
	for i := k; i < len(s); i++ {
		// removing the first element from current subarray
		if vowels[i-k] {
			count--
		}
		// adding current element to end of subarray
		if vowels[i] {
			count++
		}
		res = max(res, count)
	}
	return res
}

func isVowel(ch rune) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
