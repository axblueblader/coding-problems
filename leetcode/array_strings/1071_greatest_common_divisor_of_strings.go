package leetcode

import "strings"

/*
* Naive approach
* Iterate both string, if there is a difference then there's no common divisor
* ABAB
* AC
*
* ABABAB
* AB
* The smaller one will be the common divisor if it can keep matching the longer one
*
* ABABAB
* ABAB
* We need to remember the previous longest common divisor as we check
*
* ABABAC
* AB
* After we have the divisor we need to check till the end
*
* harder to spot observations:
* notice s1 + s2 == s2 + s1 must be true
* because
* s1 = A + A + C
* s2 = A
* then AACA != AAAC
* notice the length of gcd must be gcd(len(s1),len(s2)) and is the prefix length for either s1 or s2
* ABABAB = 8
* ABAB = 4
* gcd(8,4) = 2
* s1[0:2] = AB = gcd(s1,s2)
 */
func gcdOfStrings(str1 string, str2 string) string {
	var commonDiv strings.Builder
	var result string
	i := 0
	for ; i < len(str1) && i < (len(str2)); i++ {
		if str1[i] != str2[i] {
			return ""
		} else {
			commonDiv.WriteByte(str1[i])
			if len(str1)%commonDiv.Len() == 0 && len(str2)%commonDiv.Len() == 0 {
				result = commonDiv.String()
			}
		}
	}

	combined := str1 + str2
	for j := 0; i < len(combined); i++ {
		if combined[i] != result[i] {
			return ""
		}
		if j == len(result)-1 {
			j = 0
		} else {
			j++
		}
	}

	return result
}
