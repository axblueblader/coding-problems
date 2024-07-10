package leetcode

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/

// Traverse the tree while keep tracking of the max we've seen so far
// All nodes being traversed after should check with this max
func leetcode1448() int {
	type TreeNode struct {
		Left  *TreeNode
		Right *TreeNode
		Val   int
	}
	goodNotes := func(root *TreeNode) int {
		answer := 0
		var dfs func(*TreeNode, int)
		dfs = func(node *TreeNode, maxSoFar int) {
			if node.Val >= maxSoFar {
				answer++
				maxSoFar = node.Val
			}
			if node.Left != nil {
				dfs(node.Left, maxSoFar)
			}
			if node.Right != nil {
				dfs(node.Right, maxSoFar)
			}
		}
		dfs(root, root.Val)

		return answer
	}
	return goodNotes(nil)
}
