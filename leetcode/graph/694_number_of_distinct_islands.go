package leetcode

import "fmt"

// https://leetcode.com/problems/number-of-distinct-islands/

// DFS to find the island
// Same means translated (no rotated or reflected) which means it just need to have exacly same shape
// all a[i] == b[i]
// Stores and serialize it's shape
// For each new island discovered, check it against previous shapes
// First check if they have same dimensions width, height
// Then check each cell to determine uniqueness
type Island struct {
	top   int
	bot   int
	left  int
	right int
	id    int
}

func numDistinctIslands(grid [][]int) int {
	var dfs func(int, int, *Island)
	dfs = func(x int, y int, island *Island) {
		if x >= len(grid) || x < 0 || y < 0 || y >= len(grid[x]) {
			return
		}
		if grid[x][y] == 0 || grid[x][y] > 1 {
			return
		}
		island.top = min(island.top, x)
		island.bot = max(island.bot, x)
		island.left = min(island.left, y)
		island.right = max(island.right, y)
		// mark island as visited with a unique id
		grid[x][y] = island.id
		dfs(x+1, y, island)
		dfs(x-1, y, island)
		dfs(x, y+1, island)
		dfs(x, y-1, island)
	}

	uniq := make(map[string]bool)
	serialize := func(island *Island) string {
		str := ""
		for i := island.top; i <= island.bot; i++ {
			for j := island.left; j <= island.right; j++ {
				translatedX := i - island.top
				translatedY := j - island.left
				if grid[i][j] == island.id {
					str = str + fmt.Sprintf("[%d,%d]", translatedX, translatedY)
				}
			}
		}
		return str
	}
	islandId := 1
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 || grid[x][y] > 1 {
				continue
			}
			islandId++
			island := &Island{
				top:   x,
				bot:   x,
				left:  y,
				right: y,
				id:    islandId,
			}
			dfs(x, y, island)
			// fmt.Println(island)
			uniq[serialize(island)] = true
		}
	}

	// fmt.Println(grid)
	// fmt.Println(uniq)
	return len(uniq)
}
