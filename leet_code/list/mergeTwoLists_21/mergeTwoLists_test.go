package mergeTwoLists_21

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

		var minVal int
		if list1.Val <= list2.Val {
			minVal = list1.Val
			list1 = list1.Next
		} else {
			minVal = list2.Val
			list2 = list2.Next
		}

		head.Next = &ListNode{Val: minVal}
		head = head.Next
	}

	for list1 != nil {
		head.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		head = head.Next
	}

	for list2 != nil {
		head.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		head = head.Next
	}

	return protect.Next
}
