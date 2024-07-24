package leetcode

// https://leetcode.com/problems/number-of-islands-ii/

// Naive: for each position, find all islands
// If position connected to a known island, we don't need to search again
// i can mark each island found so whenever a new position is connected to a known island we won't increase island count
// when we connect multiple island, we need to mark the previous again, joining the islands
func numIslands2(m int, n int, positions [][]int) []int {
	ans := []int{0}
	grid := make([][]int, m+2)
	for i := 0; i < m+2; i++ {
		grid[i] = make([]int, n+2)
	}

	var dfs func(int, int, int) int
	dfs = func(x int, y int, join int) int {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
			return 0
		}
		if grid[x][y] == join {
			return 0
		}

		if grid[x][y] == 0 {
			return 0
		}

		grid[x][y] = join
		dfs(x+1, y, join)
		dfs(x-1, y, join)
		dfs(x, y+1, join)
		dfs(x, y-1, join)
		return 1
	}

	for _, pos := range positions {

		x := pos[0] + 1
		y := pos[1] + 1
		if grid[x][y] != 0 {
			ans = append(ans, ans[len(ans)-1])
			continue
		}

		joinIsland := max(grid[x-1][y], grid[x+1][y], grid[x][y-1], grid[x][y+1])
		if joinIsland == 0 {
			grid[x][y] = len(ans)
			ans = append(ans, ans[len(ans)-1]+1)
		} else {
			// join all the islands
			grid[x][y] = joinIsland
			jointed := dfs(x+1, y, joinIsland) + dfs(x-1, y, joinIsland) + dfs(x, y+1, joinIsland) + dfs(x, y-1, joinIsland)
			ans = append(ans, ans[len(ans)-1]-jointed)
		}

	}

	return ans[1:]
}
