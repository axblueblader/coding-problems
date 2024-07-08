package leetcode

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Next *ListNode
	Val  int
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	current := root
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	} else if list2 == nil {
		current.Next = list1
	}
	return root.Next
}
