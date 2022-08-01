package mergeKLists_23

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

var binaryHeap []*ListNode

func IsEmpty() bool {
	return len(binaryHeap) == 0
}

func Pop() *ListNode {
	top := binaryHeap[0]

	binaryHeap[0] = binaryHeap[len(binaryHeap)-1]

	index := 0
	for {
		if index*2+1 > len(binaryHeap)-1 {
			break
		}

		leftIndex := index*2 + 1

		if index*2+2 < len(binaryHeap)-1 {
			rightIndex := index*2 + 2

			if binaryHeap[leftIndex].Val > binaryHeap[rightIndex].Val {
				leftIndex = rightIndex
			}
		}

		if binaryHeap[index].Val > binaryHeap[leftIndex].Val {
			binaryHeap[index], binaryHeap[leftIndex] = binaryHeap[leftIndex], binaryHeap[index]
			index = leftIndex
		} else {
			break
		}
	}

	binaryHeap = binaryHeap[:len(binaryHeap)-1]
	return top
}

func Push(node *ListNode) {
	binaryHeap = append(binaryHeap, node)

	index := len(binaryHeap) - 1
	for {
		if index == 0 {
			break
		}

		parentIndex := (index - 1) / 2
		if binaryHeap[index].Val < binaryHeap[parentIndex].Val {
			binaryHeap[index], binaryHeap[parentIndex] = binaryHeap[parentIndex], binaryHeap[index]
			index = parentIndex
		} else {
			break
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {

	binaryHeap = []*ListNode{}

	head := &ListNode{}
	tail := head
	for _, node := range lists {
		if node != nil {
			Push(node)
		}
	}

	for {
		if IsEmpty() {
			break
		}
		top := Pop()
		tail.Next = top
		tail = tail.Next

		if top.Next != nil {
			Push(top.Next)
		}
	}

	return head.Next
}

func TestMergeKLists(t *testing.T) {
	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 4}
	a.Next.Next = &ListNode{Val: 5}

	b := &ListNode{Val: 1}
	b.Next = &ListNode{Val: 3}
	b.Next.Next = &ListNode{Val: 4}

	c := &ListNode{Val: 2}
	c.Next = &ListNode{Val: 6}

	mergeKLists([]*ListNode{a, b, c})
}
