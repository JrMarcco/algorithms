package combine_77

import (
	"fmt"
	"testing"
)

var ans [][]int
var selected []int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	selected = []int{}

	recursive(n, 1, k)

	return ans
}

func recursive(n int, index int, k int) {
	if len(selected) > k || len(selected)+(n-index+1) < k {
		return
	}
	if index == n+1 {
		addToAns()
		return
	}

	recursive(n, index+1, k)
	selected = append(selected, index)
	recursive(n, index+1, k)
	selected = selected[:len(selected)-1]
}

func addToAns() {
	tmp := make([]int, len(selected))
	for index, val := range selected {
		tmp[index] = val
	}
	ans = append(ans, tmp)
}

func TestCombine(t *testing.T) {
	fmt.Println(
		combine(4, 2),
	)
}
