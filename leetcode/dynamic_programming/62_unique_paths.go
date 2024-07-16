package leetcode

// https://leetcode.com/problems/unique-paths/description/

func uniquePaths(m int, n int) int {
	paths := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		paths[i] = make([]int, n+1)
	}
	paths[1][0] = 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m][n]
}

// We can actually do this with 1D array
// Notice in the path, we need paths[i-1][j] + paths[i][j-1]
// instead of paths[i-1][j], we can reuse the CURRENT cell to store the result
// paths[j] = paths[j] + paths[j-1], because paths[j] is storing result of i-1,j already
func uniquePaths1D(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
