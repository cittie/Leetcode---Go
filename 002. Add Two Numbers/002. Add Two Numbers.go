package _02__Add_Two_Numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = &ListNode{}
	val := l1.Val + l2.Val
	plus_one := val / 10

	node := &ListNode{val % 10, nil}
	head := node

	l1, l2 = l1.Next, l2.Next

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + plus_one
		plus_one = val / 10

		node.Next = &ListNode{val % 10, nil}

		l1, l2, node = l1.Next, l2.Next, node.Next
	}

	if l = l1; l2 != nil {
		l = l2
	}

	for l != nil {
		val := l.Val + plus_one
		plus_one = val / 10

		node.Next = &ListNode{val % 10, nil}

		l, node = l.Next, node.Next
	}

	if plus_one == 1 {
		node.Next = &ListNode{plus_one, nil}
	}

	return head
}