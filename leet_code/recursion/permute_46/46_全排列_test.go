package permute_46

import "testing"

var n int
var ans [][]int
var permutation []int
var selected []bool

func permute(nums []int) [][]int {
	n = len(nums)
	ans = [][]int{}
	permutation = []int{}
	selected = make([]bool, n)

	recursive(nums, 0)

	return ans
}

func recursive(nums []int, position int) {
	if position == n {
		addToAns()
		return
	}

	for index, value := range nums {
		if !selected[index] {
			permutation = append(permutation, value)
			selected[index] = true

			recursive(nums, position+1)
			selected[index] = false
			permutation = permutation[:len(permutation)-1]
		}
	}
}

func addToAns() {
	tmp := make([]int, len(permutation))
	for index, val := range permutation {
		tmp[index] = val
	}
	ans = append(ans, tmp)
}

func TestPermute(t *testing.T) {
	t.Log(
		permute([]int{1, 2, 3}),
	)
}
