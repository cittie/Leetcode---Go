package leetcode

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	val := l1.Val + l2.Val
	plusOne := val / 10

	node := &ListNode{val % 10, nil}
	head := node

	l1, l2 = l1.Next, l2.Next

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + plusOne
		plusOne = val / 10

		node.Next = &ListNode{val % 10, nil}

		l1, l2, node = l1.Next, l2.Next, node.Next
	}

	if l = l1; l2 != nil {
		l = l2
	}

	for l != nil {
		val := l.Val + plusOne
		plusOne = val / 10

		node.Next = &ListNode{val % 10, nil}

		l, node = l.Next, node.Next
	}

	if plusOne == 1 {
		node.Next = &ListNode{plusOne, nil}
	}

	return head
}
