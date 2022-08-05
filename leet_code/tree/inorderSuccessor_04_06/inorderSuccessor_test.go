package inorderSuccessor_04_06

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return findSuccessor(root, p.Val)
}

func findSuccessor(node *TreeNode, val int) *TreeNode {
	var ans *TreeNode
	cur := node

	for cur != nil {
		if cur.Val == val {
			// 查找到值匹配的节点，则后继在右子树中
			if cur.Right != nil {
				ans = cur.Right
				// 递归查找左点获取最小的节点
				for ; ans.Left != nil; ans = ans.Left {
				}
			}
			break
		}

		if val < cur.Val {
			if ans == nil || ans.Val > cur.Val {
				// 记录查找路径中比指定值大的节点中最小的一个点，
				// 若子树中没有匹配的节点，
				// 则后继在查找路径中
				ans = cur
			}
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ans
}
