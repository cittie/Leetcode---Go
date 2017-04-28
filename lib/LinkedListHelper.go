package Leetcode___Go

type ListNode struct {
	Val int
	Next *ListNode
}

func constructLinkedList(s []int) *ListNode {
	node := &ListNode{}
	root := node

	for _, val := range s {
		node.Val = val
		node.Next = &ListNode{}
		node = node.Next
	}

	node = nil

	return root
}

func linkedList2Slice(node *ListNode) []int {
	s := []int{}

	for node.Next != nil {
		s = append(s, node.Val)
		node = node.Next
	}

	if node.Val != 0 {
		s = append(s, node.Val)
	}

	return s
}
