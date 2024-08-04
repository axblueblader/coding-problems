package leetcode

// https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	magFreq := make(map[rune]int)
	for _, ch := range magazine {
		magFreq[ch]++
	}
	noteFreq := make(map[rune]int)
	for _, ch := range ransomNote {
		noteFreq[ch]++
	}

	for key, val := range noteFreq {
		if magFreq[key] < val {
			return false
		}
	}
	return true
}
