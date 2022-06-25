package minDepth_111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	minLeft := minDepth(root.Left)
	minRight := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		if minLeft > minRight {
			return minLeft + 1
		}
		return minRight + 1
	}
	if minLeft > minRight {

		return minRight + 1
	}
	return minLeft + 1
}
