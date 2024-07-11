package leetcode

import "strconv"

// https://leetcode.com/problems/equal-row-and-column-pairs/description/

// We can store seen rows as a string in a map for quick checking once we have all the columns
func equalPairs(grid [][]int) int {
	n := len(grid)
	rows := make(map[string]int)
	cols := make(map[int]string)
	ans := 0
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			// Using "-" so it if the numbers are equal we still know they are from different cell
			row = row + "-" + strconv.Itoa(grid[i][j])
			cols[j] = cols[j] + "-" + strconv.Itoa(grid[i][j])
		}
		rows[row]++
	}
	// fmt.Println(rows)
	// fmt.Println(cols)
	for _, col := range cols {
		ans = ans + rows[col]
	}
	return ans
}
