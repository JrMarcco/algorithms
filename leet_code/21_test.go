package leet_code

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	protect := &ListNode{}
	head := protect

	for list1 != nil && list2 != nil {
		newNode := &ListNode{Next: nil}

		if list1.Val <= list2.Val {
			newNode.Val = list1.Val
			list1 = list1.Next
		} else {
			newNode.Val = list2.Val
			list2 = list2.Next
		}

		head.Next = newNode
		head = head.Next
	}

	for list1 != nil {
		newNode := &ListNode{Val: list1.Val, Next: nil}
		head.Next = newNode
		head = head.Next
		list1 = list1.Next
	}

	for list2 != nil {
		newNode := &ListNode{Val: list2.Val, Next: nil}
		head.Next = newNode
		head = head.Next
		list2 = list2.Next
	}

	return protect.Next
}

func TestMerge(t *testing.T) {
	list1 := &ListNode{Val: 4, Next: nil}
	list1 = &ListNode{Val: 2, Next: list1}
	list1 = &ListNode{Val: 1, Next: list1}

	list2 := &ListNode{Val: 3, Next: nil}
	list2 = &ListNode{Val: 2, Next: list2}
	list2 = &ListNode{Val: 1, Next: list2}

	list := mergeTwoLists(list1, list2)
	for list != nil {
		t.Logf("%d ", list.Val)
		list = list.Next
	}
}
