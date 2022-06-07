package stack

import "testing"

type rectangle struct {
	width  int
	height int
}

func largestRectangleArea(heights []int) int {
	// 末尾补 0 保证栈最后能弹空
	heights = append(heights, 0)

	var rectangles []rectangle
	ans := 0
	for _, height := range heights {
		accumulatedWidth := 0
		for {
			if len(rectangles) == 0 || rectangles[len(rectangles)-1].height < height {
				break
			}
			accumulatedWidth += rectangles[len(rectangles)-1].width
			area := accumulatedWidth * rectangles[len(rectangles)-1].height
			if area > ans {
				ans = area
			}
			rectangles = rectangles[:len(rectangles)-1]
		}
		rectangles = append(rectangles, rectangle{width: accumulatedWidth + 1, height: height})
	}
	return ans
}

func TestLargest(t *testing.T) {
	t.Logf("%d", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	t.Logf("%d", largestRectangleArea([]int{2, 3}))
	t.Logf("%d", largestRectangleArea([]int{2, 4}))
}
