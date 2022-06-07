package stack

import "testing"

type rect struct {
	width  int
	height int
}

func trap(heights []int) int {
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

func TestTrap(t *testing.T) {
	t.Logf("%d", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Logf("%d", trap([]int{4, 2, 0, 3, 2, 5}))
}
