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
	if len(s) == 0 {
		return nil
	}

	var root, tail *ListNode

	node := new(ListNode)
	root = node

	for _, val := range s {
		node.Val = val
		node.Next = new(ListNode)
		tail = node
		node = node.Next
	}

	tail.Next = nil

	return root
}

// LinkedList2Slice convert ListNodes to int slice
func LinkedList2Slice(node *ListNode) []int {
	s := make([]int, 0)

	for node != nil {
		s = append(s, node.Val)
		node = node.Next
	}

	return s
}
