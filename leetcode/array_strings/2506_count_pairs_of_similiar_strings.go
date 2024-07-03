package leetcode

// https://leetcode.com/problems/count-pairs-of-similar-strings/description/

// Naive approach: check each pair with each other
// Use a set to keep track of characters
// For each character in word[i], check if it exists in word[j]
// In other words get the set difference, if it's 0 then they are a pair
func similarPairs(words []string) int {
	// empty struct is more memory efficient
	chSetList := make([]map[rune]struct{}, len(words))
	for i, word := range words {
		chSet := make(map[rune]struct{})
		for _, ch := range word {
			// put an empty struct
			chSet[ch] = struct{}{}
		}
		chSetList[i] = chSet
	}
	count := 0
	for i := 0; i < len(words); i++ {
		iSet := chSetList[i]
		for j := i + 1; j < len(words); j++ {
			jSet := chSetList[j]
			if containsAll(iSet, jSet) {
				count++
			}
		}
	}
	return count
}

func containsAll(a map[rune]struct{}, b map[rune]struct{}) bool {
	big := a
	small := b
	if len(a) < len(b) {
		small = a
		big = b
	}
	for key := range big {
		_, ok := small[key]
		if !ok {
			return false
		}
	}
	return true
}
