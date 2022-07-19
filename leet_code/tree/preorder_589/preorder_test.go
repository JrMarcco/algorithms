package preorder_589

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var stack []*Node
	var seq []int

	stack = append(stack, root)
	for {
		if len(stack) == 0 {
			break
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		seq = append(seq, top.Val)

		if top.Children != nil {
			for index := len(top.Children) - 1; index >= 0; index-- {
				stack = append(stack, top.Children[index])
			}
		}
	}

	return seq
}

func TestPreorder(t *testing.T) {

}
