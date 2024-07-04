package leetcode

// https://leetcode.com/problems/daily-temperatures/

// Warmer means higher in this context
// At position i, the question is
// "Is temp[i] the larger than a temp i've seen before that haven't got answered"
// "seen before" -> store them
// if they are answered -> remove them
// can store them as smallest first because everytime there is a larger number
// it will be removed already
func dailyTemperatures(temperatures []int) []int {
	colder := []int{}
	answer := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(colder) > 0 {
			j := colder[len(colder)-1]
			if temp <= temperatures[j] {
				break
			}
			answer[j] = i - j
			colder = colder[:len(colder)-1]
		}
		colder = append(colder, i)
	}
	return answer
}
