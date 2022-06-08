package stack

import "testing"

type rect struct {
	width  int
	height int
}

func trap1(heights []int) int {
	var rects []rect
	ans := 0
	for _, height := range heights {
		accumulatedWidth := 0
		for {
			if len(rects) == 0 || rects[len(rects)-1].height >= height {
				break
			}
			bottom := rects[len(rects)-1].height
			accumulatedWidth += rects[len(rects)-1].width
			rects = rects[:len(rects)-1]

			if len(rects) != 0 {
				var up int
				if height <= rects[len(rects)-1].height {
					up = height
				} else {
					up = rects[len(rects)-1].height
				}
				ans += accumulatedWidth * (up - bottom)
			}
		}
		rects = append(rects, rect{width: accumulatedWidth + 1, height: height})
	}
	return ans
}

func trap2(heights []int) int {
	ans := 0
	n := len(heights)
	preMax := make([]int, n)
	sufMax := make([]int, n)

	// 记录最大前缀
	preMax[0] = heights[0]
	for i := 1; i < n; i++ {
		if heights[i] > preMax[i-1] {
			preMax[i] = heights[i]
		} else {
			preMax[i] = preMax[i-1]
		}
	}

	// 记录最大后缀
	sufMax[n-1] = heights[n-1]
	for i := n - 2; i >= 0; i-- {
		if heights[i] > sufMax[i+1] {
			sufMax[i] = heights[i]
		} else {
			sufMax[i] = sufMax[i+1]
		}
	}

	for index, height := range heights {
		if index == 0 || index == n-1 {
			continue
		}
		// 取最大前缀和最大后缀中的最小值 up
		var up int
		if preMax[index-1] <= sufMax[index+1] {
			up = preMax[index-1]
		} else {
			up = sufMax[index+1]
		}
		if up > height {
			// 前后缀最小比当前值高度要高则表示能接住水，答案累加
			ans += up - height
		}
	}
	return ans
}

func TestTrap(t *testing.T) {
	t.Logf("%d", trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Logf("%d", trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Logf("%d", trap1([]int{4, 2, 0, 3, 2, 5}))
	t.Logf("%d", trap2([]int{4, 2, 0, 3, 2, 5}))
}
