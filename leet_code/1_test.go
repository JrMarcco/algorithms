package leet_code

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}

	for index, value := range nums {
		if targetIndex, ok := numsMap[target-value]; ok {
			return []int{targetIndex, index}
		}
		numsMap[value] = index
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	t.Log(
		twoSum([]int{2, 7, 11, 15}, 9),
	)
	t.Log(
		twoSum([]int{3, 2, 4}, 6),
	)
}
