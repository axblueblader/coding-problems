package leetcode

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// if we split down the middle, the twin node are nodes with equal amount of distance to middle
// we can traverse whole list, put into array, start from middle and move left right at same step to get twin sum
// N is even
// fast and slow pointer
// [1,2,3,4,5,6]
//
//	s     f
//
// when at middle try to reverse the left side
func pairSum(head *ListNode) int {
	slow := head
	fast := head.Next
	var next, prev *ListNode
	for fast != nil {
		// reverse the list
		next = slow.Next
		slow.Next = prev
		prev = slow
		slow = next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}

	}
	left := prev
	right := next
	ans := 0
	for right != nil {
		if left.Val+right.Val > ans {
			ans = left.Val + right.Val
		}

		left = left.Next
		right = right.Next
	}
	return ans
}
