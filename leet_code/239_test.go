package leet_code

import (
	"fmt"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	var ans, deque []int

	for index, num := range nums {
		// 保证队头合法性
		for len(deque) > 0 && deque[0] <= index-k {
			deque = deque[1:]
		}
		// 维护队列单调性，插入新的选项
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, index)
		// 取队头更新答案
		if index >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}

func TestMax(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
