package leetcode

// https://leetcode.com/problems/regions-cut-by-slashes/description/

// a slash can divide a square into 2 regions
// Intuition is to use BFS
// But how to check the path and slashes
// |001| |100|
// |010| |010|
// |100| |001|

// Up scale the board to 3x3
// Then we can BFS as usual

func regionsBySlashes(grid []string) int {
	n := len(grid)
	scaled := make([][]int, n*3)
	for i := 0; i < n*3; i++ {
		scaled[i] = make([]int, n*3)
	}
	for x := range grid {
		for y := range grid[x] {
			switch grid[x][y] {
			case '/':
				scaled[x*3][y*3+2] = 1
				scaled[x*3+1][y*3+1] = 1
				scaled[x*3+2][y*3] = 1
			case '\\':
				scaled[x*3][y*3] = 1
				scaled[x*3+1][y*3+1] = 1
				scaled[x*3+2][y*3+2] = 1
			}
		}
	}
	n3 := n * 3
	count := 0
	bfs := func(x int, y int) {
		queue := [][]int{{x, y}}
		for len(queue) > 0 {
			levelSize := len(queue)
			for i := 0; i < levelSize; i++ {
				pos := queue[0]
				queue = queue[1:]
				if pos[0] < 0 || pos[1] < 0 || pos[0] >= n3 || pos[1] >= n3 || scaled[pos[0]][pos[1]] == 1 {
					continue
				}
				scaled[pos[0]][pos[1]] = 1
				queue = append(queue,
					[]int{pos[0] + 1, pos[1]},
					[]int{pos[0] - 1, pos[1]},
					[]int{pos[0], pos[1] + 1},
					[]int{pos[0], pos[1] - 1})
			}
		}
	}

	for x := range scaled {
		for y := range scaled[x] {
			if scaled[x][y] == 0 {
				// BFS
				bfs(x, y)
				count++
			}
		}
	}

	return count
}
