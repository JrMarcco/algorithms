package maxDepth_104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)

	if maxLeft > maxRight {
		return maxLeft + 1
	}
	return maxRight + 1
}
