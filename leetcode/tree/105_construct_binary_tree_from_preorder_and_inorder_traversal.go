package leetcode

import "slices"

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder means root, left, right
// inorder means left, root, right
// using preorder array, we can immediately get the current
// if we know the position of the root note in the inorder array
// everything on the left will be it's left subtree
// elements on the right will belong to it's right subtree
// knowing the number of elements, we'll know which subarray belongs to the left subtree in the inorder array
// recursively splitting the array up with these steps, we can build the correct tree
func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func([]int, []int) *TreeNode
	build = func(pre []int, ino []int) *TreeNode {
		if len(pre) == 0 {
			return nil
		}
		posInorder := slices.Index(ino, pre[0])

		preRightSubTree := pre[posInorder+1:]
		preLeftSubTree := pre[1 : posInorder+1]
		inoLeftSubTree := ino[:posInorder]
		inoRightSubTree := ino[posInorder+1:]

		return &TreeNode{
			Val:   pre[0],
			Left:  build(preLeftSubTree, inoLeftSubTree),
			Right: build(preRightSubTree, inoRightSubTree),
		}
	}

	return build(preorder, inorder)
}
