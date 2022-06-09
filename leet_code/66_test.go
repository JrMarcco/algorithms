package leet_code

import (
	"fmt"
	"testing"
)

func plusOne(digits []int) []int {

	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			continue
		}
		digits[i]++
		// 处理中间不为 9 尾数全为 9 的情况
		// 例如 1 5 9 9 9
		// 将 5 -> 6 尾数均置为 0
		for j := i + 1; j < n; j++ {
			digits[j] = 0
		}
		return digits
	}

	// 所有数字都为 9 的情况
	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}

func TestPlusOne(t *testing.T) {
	fmt.Println(plusOne([]int{1, 4, 8}))
}
