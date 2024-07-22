package leetcode

// https://leetcode.com/problems/reverse-nodes-in-k-group/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var prev *ListNode
	current := head
	groupTail := head
	dummy := &ListNode{}
	prevGroupTail := dummy
	count := 1
	for current != nil {
		next := current.Next
		current.Next = prev

		if count%k == 0 {
			prev = groupTail
			prevGroupTail.Next = current
			prevGroupTail = groupTail
			groupTail = next
		} else {
			prev = current
		}
		count++

		current = next
	}
	prevGroupTail.Next = groupTail
	current = prev
	prev = nil
	count--
	for count%k != 0 {
		next := current.Next
		current.Next = prev

		prev = current
		current = next

		count--
	}
	return dummy.Next
}
