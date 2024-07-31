package leetcode

// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/description/

import "slices"

func nearestExit(maze [][]byte, entrance []int) int {
	// exit is a zero byte
	exit := byte(0)
	n := len(maze)
	// pad the maze
	pad := make([]byte, n+2)
	// pad top and bottom
	maze = append(maze, pad)
	maze = append([][]byte{pad}, maze...)
	// pad left and right
	for row := range maze {
		maze[row] = append(maze[row], byte(0))
		maze[row] = append([]byte{0}, maze[row]...)
	}

	entrance[0] += 1
	entrance[1] += 1
	queue := [][]int{entrance}
	// fmt.Println(maze, queue)
	stepCount := 0

	// use BFS to find shortest path
	for len(queue) > 0 {

		size := len(queue)

		for i := 0; i < size; i++ {
			pos := queue[0]
			queue = queue[1:]
			if maze[pos[0]][pos[1]] == '+' {
				continue
			}
			up := maze[pos[0]-1][pos[1]]
			down := maze[pos[0]+1][pos[1]]
			left := maze[pos[0]][pos[1]-1]
			right := maze[pos[0]][pos[1]+1]

			if !slices.Equal(pos, entrance) && (up == exit || down == exit || left == exit || right == exit) {
				return stepCount
			}
			// mark as visited
			maze[pos[0]][pos[1]] = '+'
			if up == '.' {
				queue = append(queue, []int{pos[0] - 1, pos[1]})
			}
			if down == '.' {
				queue = append(queue, []int{pos[0] + 1, pos[1]})
			}
			if left == '.' {
				queue = append(queue, []int{pos[0], pos[1] - 1})
			}
			if right == '.' {
				queue = append(queue, []int{pos[0], pos[1] + 1})
			}
		}

		stepCount++
	}

	return -1
}
