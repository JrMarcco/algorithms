package reverseList_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var last *ListNode

	for {
		if head == nil {
			break
		}

		next := head.Next
		head.Next = last
		last = head
		head = next
	}

	return last
}
