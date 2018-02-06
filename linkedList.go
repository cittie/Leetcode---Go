package leetcode

// ListNode for singly-linked list
/* type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// ConstructLinkedList convert int slice to ListNodes and return the head
func ConstructLinkedList(s []int) *ListNode {
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

// LinkedList2Slice convert ListNodes to int slice
func LinkedList2Slice(node *ListNode) []int {
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
