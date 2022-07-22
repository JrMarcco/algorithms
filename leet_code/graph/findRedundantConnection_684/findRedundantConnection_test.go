package findRedundantConnection_684

import "testing"

func findRedundantConnection(edges [][]int) []int {
	n := 0
	for _, edge := range edges {
		n = max(n, max(edge[0], edge[1]))
	}

	edgeArray := make([][]int, n+1)
	visitArray := make([]bool, n+1)
	hasCycle := false

	addEdge := func(u int, v int) {
		edgeArray[u] = append(edgeArray[u], v)
	}

	var dfs func(x int, father int)
	dfs = func(x int, father int) {
		visitArray[x] = true

		for _, y := range edgeArray[x] {
			if y == father {
				continue
			}
			if visitArray[y] {
				hasCycle = true
				return
			}
			dfs(y, x)
		}
		visitArray[x] = false
	}

	for _, edge := range edges {

		addEdge(edge[0], edge[1])
		addEdge(edge[1], edge[0])

		dfs(edge[0], -1)

		if hasCycle {
			return edge
		}
	}

	return []int{}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestFindRedundantConnection(t *testing.T) {
	t.Log(
		findRedundantConnection([][]int{
			{1, 2}, {1, 3}, {2, 3},
		}),
	)

	t.Log(
		findRedundantConnection([][]int{
			{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5},
		}),
	)
}
