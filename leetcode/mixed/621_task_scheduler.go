package leetcode

// https://leetcode.com/problems/task-scheduler/

// Greedy, on each iteration, schedule the task with most remaining occurrence
// because we want to spread out the tasks, to avoid waiting on cooldowns
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	counts := make(map[byte]int)
	cd := make(map[byte]int)
	for _, v := range tasks {
		counts[v]++
	}
	end := len(tasks)
	step := 0
	for done := 0; done < end; step++ {
		var maxK byte
		maxV := 0
		for k, v := range counts {
			if cd[k] <= step && v > maxV {
				maxK = k
				maxV = v
			}
		}
		if maxK == 0 {
			continue
		}
		cd[maxK] = step + n + 1
		counts[maxK]--
		done++
	}
	return step
}
