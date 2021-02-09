package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 思路：双指针
	// p1先移动n步，然后p1和p2一起移动，直到尾部
	// 此时p2就是要删的节点

	pre := &ListNode{Next: head}

	p1, p2 := pre, pre

	for ; n > 0; n-- {
		p1 = p1.Next
	}

	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return pre.Next
}
