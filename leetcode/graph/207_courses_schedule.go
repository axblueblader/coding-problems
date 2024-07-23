package leetcode

// https://leetcode.com/problems/course-schedule/description/

// prerequisites is like a list of edges of a directed graph
// can finish means there are no cycle in the graph
// use DFS on each course and check for cycle
// we store an extra 'can learn' map to denotes the courses we know that are possible to learn to reduce DFS steps
func canFinish(numCourses int, prerequisites [][]int) bool {
	toLearn := make(map[int]bool)
	canLearn := make(map[int]bool)
	graph := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, numCourses)
	}
	for _, pair := range prerequisites {
		graph[pair[0]][pair[1]] = 1
	}

	var dfs func(int, map[int]bool) bool
	dfs = func(course int, toLearn map[int]bool) bool {
		if canLearn[course] {
			return false
		}
		if toLearn[course] == true {
			return true
		}

		toLearn[course] = true
		for i := 0; i < numCourses; i++ {
			if graph[course][i] != 0 {
				cyclic := dfs(i, toLearn)
				if cyclic {
					return true
				}
			}
		}
		toLearn[course] = false

		return false
	}

	for i := 0; i < numCourses; i++ {
		cyclic := dfs(i, toLearn)
		if cyclic {
			return false
		}
		canLearn[i] = true
	}

	return true
}
