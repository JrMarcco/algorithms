package findOrder_210

import "testing"

func findOrder(numCourses int, prerequisites [][]int) []int {
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
	if len(finished) == numCourses {
		return finished
	}

	return []int{}
}

func TestFindOrder(t *testing.T) {
	t.Log(
		findOrder(3, [][]int{{1, 0}, {1, 2}, {0, 1}}),
	)

	t.Log(
		findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}),
	)
}
