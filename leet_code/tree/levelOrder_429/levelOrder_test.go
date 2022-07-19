package levelOrder_429

type Node struct {
	Val      int
	Children []*Node
}

type NodeLevel struct {
	node  *Node
	level int
}

func levelOrder(root *Node) [][]int {

	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	var queue []NodeLevel

	queue = append(queue, NodeLevel{root, 0})

	currentLevel := 0
	var subAns []int

	for {
		if len(queue) == 0 {
			break
		}

		front := queue[0]
		queue = queue[1:]

		if front.node.Children != nil {
			for _, child := range front.node.Children {
				queue = append(queue, NodeLevel{child, front.level + 1})
			}
		}

		if front.level != currentLevel {
			currentLevel = front.level
			ans = append(ans, subAns)
			subAns = []int{}
		}
		subAns = append(subAns, front.node.Val)
	}
	ans = append(ans, subAns)
	return ans
}
