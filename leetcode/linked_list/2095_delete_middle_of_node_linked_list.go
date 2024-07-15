package leetcode

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// keep track of middle node as we traverse
func deleteMiddle(head *ListNode) *ListNode {
	nodeCount := 0
	keepHead := head
	middleNode := head
	currentNode := head
	midIdx := 0
	for currentNode != nil {
		nodeCount++
		for midIdx < (nodeCount/2)-1 {
			midIdx++
			middleNode = middleNode.Next
		}
		currentNode = currentNode.Next
	}
	if middleNode.Next != nil {
		middleNode.Next = middleNode.Next.Next
	} else {
		return nil
	}
	return keepHead
}
