package inorderTraversal_94

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var seq []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	seq = []int{}
	dfs(root)
	return seq
}

func dfs(root *TreeNode) {
	if root.Left != nil {
		dfs(root.Left)
	}
	seq = append(seq, root.Val)
	if root.Right != nil {
		dfs(root.Right)
	}
}

func TestInorderTraversal(t *testing.T) {

}
