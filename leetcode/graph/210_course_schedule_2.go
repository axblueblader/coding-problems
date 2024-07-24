package leetcode

// https://leetcode.com/problems/course-schedule-ii/description/

// using topological sort with DFS to print order of dependencies
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = []int{}
	}

	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	visited := make([]bool, numCourses)
	learned := make([]bool, numCourses)
	answer := []int{}
	var topoSort func(int) bool
	topoSort = func(course int) bool {
		visited[course] = true

		for _, prereq := range graph[course] {
			if !visited[prereq] && !topoSort(prereq) {
				// haven't visit and topo sort found cycle
				return false
			} else if !learned[prereq] {
				// visited but haven't learned means there's a cyclic dependency
				return false
			}
		}

		learned[course] = true
		answer = append(answer, course)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] && !topoSort(i) {
			return []int{}
		}
	}

	return answer
}
