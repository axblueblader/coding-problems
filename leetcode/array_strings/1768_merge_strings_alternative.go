package leetcode

import "strings"

// Each iteration we can append word1[i] then word[2]
// then append the rest based on whichever is longer
// some small optimization involves using the better string manipulation method
func mergeAlternately(word1 string, word2 string) string {
	var result strings.Builder
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		result.WriteString(string(word1[i]))
		result.WriteString(string(word2[i]))
	}
	for ; i < len(word1); i++ {
		result.WriteString(string(word1[i]))
	}
	for ; i < len(word2); i++ {
		result.WriteString(string(word2[i]))
	}
	return result.String()
}
