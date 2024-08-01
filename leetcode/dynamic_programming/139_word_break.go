package leetcode

// https://leetcode.com/problems/word-break/

// leet code
// leetc - leetco - leetcode
// [leet, leetcode]
// backtracking

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// normal DFS got timed out, so need to memoize
	unmatchable := make([]bool, len(s))
	var matchWord func(int, int) bool
	matchWord = func(start int, end int) bool {
		if unmatchable[start] {
			// fmt.Println("cache hit", start, end+1)
			return false
		}
		for ; end < len(s); end++ {
			// fmt.Println(s[start : end+1])
			if !wordMap[s[start:end+1]] {
				continue
			}

			// this is the final word matched in the string
			if end+1 == len(s) {
				return true
			}

			// do again with next word starting from after last character
			if matchWord(end+1, end+1) {
				return true
			} else {
				unmatchable[end+1] = true
			}
		}
		return false
	}

	return matchWord(0, 0)
}
