package canFinish_207

import "testing"

func canFinish(numCourses int, prerequisites [][]int) bool {
	edgeArray := make([][]int, numCourses)
	inDegreeArray := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		edgeArray[prerequisite[1]] = append(edgeArray[prerequisite[1]], prerequisite[0])
		inDegreeArray[prerequisite[0]]++
	}

	var queue []int
	for index, degree := range inDegreeArray {
		if degree == 0 {
			queue = append(queue, index)
		}
	}

	var finished []int
	for {
		if len(queue) == 0 {
			break
		}
		front := queue[0]
		queue = queue[1:]
		finished = append(finished, front)

		for _, course := range edgeArray[front] {
			inDegreeArray[course]--
			if inDegreeArray[course] == 0 {
				queue = append(queue, course)
			}
		}
	}

	return len(finished) == numCourses
}

func TestCanFinish(t *testing.T) {
	t.Log(
		canFinish(2, [][]int{{1, 0}}),
	)
}
