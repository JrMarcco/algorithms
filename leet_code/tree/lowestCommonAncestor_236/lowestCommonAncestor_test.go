package lowestCommonAncestor_236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeRes struct {
	Left  bool
	Right bool
}

var LeftNode *TreeNode
var RightNode *TreeNode
var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	LeftNode = p
	RightNode = q
	ans = nil

	dfs(root)

	return ans
}

func dfs(root *TreeNode) NodeRes {

	if root == nil {
		return NodeRes{Left: false, Right: false}
	}

	leftResult := dfs(root.Left)
	rightResult := dfs(root.Right)
	result := NodeRes{}

	result.Left = leftResult.Left || rightResult.Left || root.Val == LeftNode.Val
	result.Right = leftResult.Right || rightResult.Right || root.Val == RightNode.Val

	if result.Left && result.Right && ans == nil {
		ans = root
	}
	return result
}
