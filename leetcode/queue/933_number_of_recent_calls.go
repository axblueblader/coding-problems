package leetcode

// https://leetcode.com/problems/number-of-recent-calls/description/

type RecentCounter struct {
	Queue []int
	Range int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Range: 3000,
		Queue: []int{},
	}
}

func (recentCounter *RecentCounter) Ping(t int) int {
	threshold := t - recentCounter.Range
	for len(recentCounter.Queue) > 0 && recentCounter.Queue[0] < threshold {
		// remove first element
		recentCounter.Queue = recentCounter.Queue[1:]
	}
	recentCounter.Queue = append(recentCounter.Queue, t)
	return len(recentCounter.Queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
