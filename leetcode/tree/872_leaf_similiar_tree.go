package leetcode

// https://leetcode.com/problems/leaf-similar-trees/

func leetcode872() bool {
	type TreeNode struct {
		Left  *TreeNode
		Right *TreeNode
		Val   int
	}

	leafSimilar := func(root1 *TreeNode, root2 *TreeNode) bool {
		sequence1 := []int{}
		sequence2 := []int{}

		var dfs func(*TreeNode, *[]int)
		dfs = func(node *TreeNode, sequence *[]int) {
			if node.Left == nil && node.Right == nil {
				*sequence = append(*sequence, node.Val)
				return
			}
			if node.Left != nil {
				dfs(node.Left, sequence)
			}
			if node.Right != nil {
				dfs(node.Right, sequence)
			}
		}
		dfs(root1, &sequence1)
		dfs(root2, &sequence2)
		if len(sequence1) != len(sequence2) {
			return false
		}
		for i := range sequence1 {
			if sequence1[i] != sequence2[i] {
				return false
			}
		}
		return true
	}
	return leafSimilar(nil, nil)
}
