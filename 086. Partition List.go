package leetcode

func partition(head *ListNode, x int) *ListNode {
	var left, right, root, tail *ListNode
	left, right = new(ListNode), new(ListNode)
	root, tail = left, right

	for ; head != nil; head = head.Next {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
	}

	left.Next, right.Next = tail.Next, nil

	return root.Next
}
