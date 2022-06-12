package leet_code

import (
	"testing"
)

func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := map[int64]int{}

	if len(obstacles) != 0 {
		for _, point := range obstacles {
			obstaclesMap[getHashKey(point[0], point[1])] = 1
		}
	}

	x, y, ans, direction := 0, 0, 0, 0

	directionX := [...]int{0, 1, 0, -1}
	directionY := [...]int{1, 0, -1, 0}

	for _, command := range commands {
		if command == -1 {
			direction = (direction + 1) % 4
			continue
		}
		if command == -2 {
			direction = (direction + 3) % 4
			continue
		}

		for i := 0; i < command; i++ {
			newX, newY := x+directionX[direction], y+directionY[direction]
			if _, ok := obstaclesMap[getHashKey(newX, newY)]; ok {
				break
			}
			x, y = newX, newY
		}
		newAns := x*x + y*y
		if newAns > ans {
			ans = newAns
		}
	}
	return ans
}

func getHashKey(x int, y int) int64 {
	return int64(
		(3000+x)*60001 + (3000 + y),
	)
}

func TestRobotSim(t *testing.T) {

	//t.Log(
	//	robotSim([]int{4, -1, 3}, [][]int{}),
	//)

	t.Log(
		robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}),
	)
}
