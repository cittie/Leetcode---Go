package _02__Add_Two_Numbers

import (
	"testing"
	"reflect"
)

func Test_addTwoNumbers(t *testing.T)  {
	first := [][]int {
		[]int {2, 4, 3},
		[]int {9, 9, 9},
		[]int {3, 7, 9},
		[]int {1},
		[]int {1, 8},
	}
	second := [][]int {
		[]int {5, 6, 4},
		[]int {9, 9, 9},
		[]int {7, 2},
		[]int {9, 9, 9, 9},
		[]int {0},
	}
	expected := [][]int {
		[]int {7, 0, 8},
		[]int {8, 9, 9, 1},
		[]int {0, 0, 0, 1},
		[]int {0, 0, 0, 0, 1},
		[]int {1, 8},
	}

	for i := range first {
		p1, p2 := constructLinkedList(first[i]), constructLinkedList(second[i])
		p3 := addTwoNumbers(p1, p2)

		r := linkedList2Slice(p3)

		if !reflect.DeepEqual(r, expected[i]) {

			t.Error("Expected: ", expected[i], "got: ", r)
		}
	}
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