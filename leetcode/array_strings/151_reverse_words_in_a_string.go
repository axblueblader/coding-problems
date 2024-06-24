package leetcode

// We could use string helper methods to Trim but i tried to use no helpers for good exercise
// The simple logic is remember a word while traversing and add it to the begining of result when a 'space' is met
//
// Small optimization, we can keep track of the index and use slicing instead of concat for the 'word'
func reverseWords(s string) string {
	word := ""
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else if word != "" {
			result = word + " " + result
			word = ""
		}
	}
	// add in the last word
	if word != "" {
		result = word + " " + result
	}
	// exclude the spare 'space' at the end
	return result[:len(result)-1]
}
