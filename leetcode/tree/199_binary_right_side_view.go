package leetcode

// https://leetcode.com/problems/binary-tree-right-side-view/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	nextQ := []*TreeNode{}
	ans := []int{}
	for len(queue) > 0 {
		// pop queue
		node := queue[0]
		queue = queue[1:]

		// add children to next level
		if node.Left != nil {
			nextQ = append(nextQ, node.Left)
		}
		if node.Right != nil {
			nextQ = append(nextQ, node.Right)
		}

		if len(queue) == 0 {
			// we've finished traversing a level, it means this is the rightmost node
			ans = append(ans, node.Val)
			// move on to the next level
			queue = nextQ
			nextQ = []*TreeNode{}
		}
	}
	return ans
}
