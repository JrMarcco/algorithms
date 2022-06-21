package subsets_78

import (
	"fmt"
	"testing"
)

var n int
var ans [][]int
var selected []int

func subsets(nums []int) [][]int {
	n = len(nums)
	ans = [][]int{}
	selected = []int{}

	recursive(nums, 0)

	return ans
}

func recursive(nums []int, index int) {
	if index == n {
		addToAns()
		return
	}

	recursive(nums, index+1)
	selected = append(selected, nums[index])
	recursive(nums, index+1)
	selected = selected[:len(selected)-1]
}

func addToAns() {
	tmp := make([]int, len(selected))
	for index, val := range selected {
		tmp[index] = val
	}
	ans = append(ans, tmp)
}

func TestSubSets(t *testing.T) {
	fmt.Println(
		subsets([]int{9, 0, 3, 5, 7}),
	)
}
