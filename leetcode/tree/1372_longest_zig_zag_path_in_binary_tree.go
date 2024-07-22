package leetcode

// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	maxVal := 0
	var dfs func(*TreeNode, int, rune)
	dfs = func(node *TreeNode, sum int, direction rune) {
		if node == nil {
			return
		}

		sum = sum + 1
		if sum > maxVal {
			maxVal = sum
		}

		if direction == 'L' {
			dfs(node.Right, sum, 'R')
			dfs(node.Left, 1, 'L')
		} else {
			dfs(node.Right, 1, 'R')
			dfs(node.Left, sum, 'L')
		}
	}

	dfs(root, 0, 'L')

	return maxVal - 1
}
