package leetcode

// https://leetcode.com/problems/odd-even-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	ansHead := &ListNode{}
	evenHead := &ListNode{}
	curr := head
	curOdd := ansHead
	curEven := evenHead
	i := 1
	for curr != nil {
		// fmt.Println(curr, curOdd, curEven)
		if i%2 != 0 {
			curOdd.Next = curr
			curOdd = curOdd.Next
		} else {
			curEven.Next = curr
			curEven = curEven.Next
		}
		i++
		curr = curr.Next
	}
	curOdd.Next = evenHead.Next
	curEven.Next = nil
	// curr = ansHead.Next
	// for curr != nil {
	//     fmt.Print(curr)
	//     curr = curr.Next
	// }
	return ansHead.Next
}
