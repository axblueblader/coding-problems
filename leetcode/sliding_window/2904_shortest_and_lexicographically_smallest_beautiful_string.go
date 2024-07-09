package leetcode

// https://leetcode.com/problems/shortest-and-lexicographically-smallest-beautiful-string/

import "sort"

// use a slding window to get the substring having k 1s
func shortestBeautifulSubstring(s string, k int) string {
	candidates := []string{}
	minLength := len(s)
	start := 0
	oneCount := 0
	for i := 0; i < len(s); i++ {
		// fmt.Println(i, start, oneCount)
		if s[i] == '0' {
			continue
		}
		for s[start] != '1' {
			start++
		}
		oneCount++
		if k-oneCount > 0 {
			continue
		}
		cur := s[start : i+1]
		if len(cur) <= minLength {
			minLength = len(cur)
			candidates = append(candidates, cur)
		}
		start++
		oneCount--
	}
	sort.Strings(candidates)
	for _, str := range candidates {
		if len(str) == minLength {
			return str
		}
	}
	return ""
}
