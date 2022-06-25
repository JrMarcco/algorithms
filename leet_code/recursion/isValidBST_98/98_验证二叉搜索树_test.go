package isValidBST_98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt32, math.MaxInt32)
}

func validate(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val < min || root.Val > max {
		return false
	}

	return validate(root.Left, min, root.Val-1) && validate(root.Right, root.Val+1, max)
}
