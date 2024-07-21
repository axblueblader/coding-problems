package leetcode

// https://leetcode.com/problems/path-sum-iii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Naive way, pass on the sums from previous nodes
func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(*TreeNode, []int)
	count := 0
	dfs = func(node *TreeNode, sums []int) {
		if node == nil {
			return
		}

		newSums := []int{0}
		for _, sum := range sums {
			if sum+node.Val == targetSum {
				count++
			}
			newSums = append(newSums, sum+node.Val)
		}
		dfs(node.Left, newSums)
		dfs(node.Right, newSums)
	}
	dfs(root, []int{0})
	return count
}

// Better way: use a prefix sum hash, storing the sum seen before and the number paths that can lead to that
// => https://www.geeksforgeeks.org/find-if-there-is-a-subarray-with-0-sum/
func traversal(root *TreeNode, targetSum int, prefix int, prefixSum map[int]int) int {
	if root == nil {
		return 0
	}

	prefix += root.Val
	count := prefixSum[prefix-targetSum]
	prefixSum[prefix]++

	leftCount := traversal(root.Left, targetSum, prefix, prefixSum)
	rightCount := traversal(root.Right, targetSum, prefix, prefixSum)

	prefixSum[prefix]--

	return count + leftCount + rightCount
}

func pathSumBetter(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	prefixSum := make(map[int]int)
	prefixSum[0] = 1

	return traversal(root, targetSum, 0, prefixSum)
}
