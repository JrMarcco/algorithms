package leet_code

import (
	"testing"
)

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

func maximalRectangle(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	newMatrix := make([][]int, len(matrix))

	// 将 0 1 二维数组逐行转换成宽度为 1 的柱状图
	// 要注意这里的输入是 char 不能直接当成 byte 处理
	// 处理第一行，当前为 ’1‘ 则高度为 1
	newMatrix[0] = make([]int, len(matrix[0]))
	for index, value := range matrix[0] {
		if value == '0' {
			newMatrix[0][index] = 0
		} else {
			newMatrix[0][index] = 1
		}
	}

	// 从第二行开始逐行累加
	for i := 1; i < len(matrix); i++ {
		newMatrix[i] = make([]int, len(matrix[i]))
		for j, value := range matrix[i] {
			// 当前为 '0' 表示柱子不连续，则在这一行的累加高度为 0
			if value == '0' {
				newMatrix[i][j] = 0
			} else {
				// 当前为 '1' 则这一行的高度为上一行的高度 +1
				newMatrix[i][j] = newMatrix[i-1][j] + 1
			}
		}
	}

	ans := 0
	for _, row := range newMatrix {
		// 逐行使用 84 题的方式求出最大面积
		tmp := largestRectangleArea(row)
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}

func TestLargest(t *testing.T) {
	//t.Logf("%d", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	//t.Logf("%d", largestRectangleArea([]int{2, 3}))
	//t.Logf("%d", largestRectangleArea([]int{2, 4}))

	t.Log(
		maximalRectangle([][]byte{{'1', '0', '1', '0', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}),
	)

	t.Log(
		maximalRectangle([][]byte{}),
	)

	t.Log(
		maximalRectangle([][]byte{{1, 0, 1, 1, 1, 0}}),
	)

	t.Log(
		maximalRectangle([][]byte{{1, 0, 1, 0, 1, 0}}),
	)
}
