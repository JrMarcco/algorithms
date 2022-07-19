package findShortestSubArray_697

import (
	"math"
	"testing"
)

type Item struct {
	count      int
	firstIndex int
	lastIndex  int
}

func findShortestSubArray(nums []int) int {

	if nums == nil {
		return 0
	}

	itemMap := map[int]Item{}
	for index, value := range nums {
		if item, ok := itemMap[value]; ok {
			item.count++
			item.lastIndex = index
			itemMap[value] = item
		} else {
			itemMap[value] = Item{
				count:      1,
				firstIndex: index,
				lastIndex:  index,
			}
		}
	}

	ans := math.MaxInt
	maxCount := 0
	for _, item := range itemMap {
		if maxCount < item.count {
			maxCount = item.count
			ans = item.lastIndex - item.firstIndex
		}

		if maxCount == item.count {
			tmp := item.lastIndex - item.firstIndex
			if tmp < ans {
				ans = tmp
			}
		}
	}

	return ans + 1
}

func TestFindShortestSubArray(t *testing.T) {
	t.Log(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	t.Log(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
