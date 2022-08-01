package maxSlidingWindow_239

import "testing"

type Pair struct {
	Val   int
	Index int
}

var binaryHeap []*Pair

func Top() *Pair {
	return binaryHeap[0]
}

func Pop() *Pair {
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

			if binaryHeap[leftIndex].Val < binaryHeap[rightIndex].Val {
				leftIndex = rightIndex
			}
		}

		if binaryHeap[index].Val < binaryHeap[leftIndex].Val {
			binaryHeap[index], binaryHeap[leftIndex] = binaryHeap[leftIndex], binaryHeap[index]
			index = leftIndex
		} else {
			break
		}
	}

	binaryHeap = binaryHeap[:len(binaryHeap)-1]
	return top
}

func Push(node *Pair) {
	binaryHeap = append(binaryHeap, node)

	index := len(binaryHeap) - 1
	for {
		if index == 0 {
			break
		}

		parentIndex := (index - 1) / 2
		if binaryHeap[index].Val > binaryHeap[parentIndex].Val {
			binaryHeap[index], binaryHeap[parentIndex] = binaryHeap[parentIndex], binaryHeap[index]
			index = parentIndex
		} else {
			break
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	binaryHeap = []*Pair{}

	var ans []int
	for index, num := range nums {
		Push(&Pair{Val: num, Index: index})

		if index >= k-1 {
			for {
				if Top().Index <= index-k {
					Pop()
					continue
				}
				break
			}
			ans = append(ans, Top().Val)
		}
	}
	return ans
}

func TestMaxSlidingWindow(t *testing.T) {
	t.Log(
		maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)
}
