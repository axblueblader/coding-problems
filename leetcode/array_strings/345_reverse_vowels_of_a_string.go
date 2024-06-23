package leetcode

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'i': true, 'e': true, 'o': true, 'u': true,
		'A': true, 'I': true, 'E': true, 'O': true, 'U': true,
	}
	res := []byte(s)
	for start, end := 0, len(s)-1; start < end; {
		if vowels[s[start]] && vowels[s[end]] {
			res[end] = s[start]
			res[start] = s[end]
			start++
			end--
			continue
		}
		if !vowels[s[start]] {
			start++
		}
		if !vowels[s[end]] {
			end--
		}
	}
	return string(res)
}

// A better way to do vowel check
func isVowel(char byte) bool {
	switch char {
	case 'a', 'A', 'i', 'I', 'e', 'E', 'o', 'O', 'u', 'U':
		return true
	}
	return false
}
