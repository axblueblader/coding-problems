package leetcode

// https://leetcode.com/problems/number-of-equivalent-domino-pairs/

type Tuple struct {
	l int
	r int
}

func numEquivDominoPairs(dominoes [][]int) int {
	counter := make(map[Tuple]int)

	for _, dom := range dominoes {
		l, r := dom[0], dom[1]
		if dom[0] > dom[1] {
			l = dom[1]
			r = dom[0]
		}
		counter[Tuple{l: l, r: r}]++
	}
	count := 0
	for _, val := range counter {
		if val > 1 {
			count += (val * (val - 1)) / 2
		}
	}
	return count
}
