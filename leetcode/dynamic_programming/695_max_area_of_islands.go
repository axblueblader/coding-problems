package leetcode

// https://leetcode.com/problems/max-area-of-island/

// how to identify island
// we can start from the left to right, top to down
// meet a 1, it means we spotted an island
// we need to explore it by going right and down
// this is similiar to DFS
// mark it as visisted already, we can do that by turning it into water, instead of using another 2D matrix to keep track
func maxAreaOfIsland(grid [][]int) int {
	var dfs func(int, int) int
	dfs = func(x int, y int) int {
		if x >= len(grid) || x < 0 {
			return 0
		}
		if y >= len(grid[x]) || y < 0 {
			return 0
		}
		if grid[x][y] == 0 {
			return 0
		}

		// update to 0 so we never visit again
		grid[x][y] = 0
		// spread to all 4 directions
		return 1 + dfs(x+1, y) + dfs(x, y+1) + dfs(x-1, y) + dfs(x, y-1)
	}

	maxArea := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				// DFS
				area := dfs(i, j)
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}
