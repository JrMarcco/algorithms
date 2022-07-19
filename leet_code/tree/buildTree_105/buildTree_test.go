package buildTree_105

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre []int
var in []int

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre = preorder
	in = inorder
	return build(0, len(pre)-1, 0, len(in)-1)
}

func build(preLeft int, preRight int, inLeft int, inRight int) *TreeNode {
	if preLeft > preRight || inLeft > inRight {
		return nil
	}

	root := &TreeNode{Val: pre[preLeft]}
	rootIndex := 0 // root 在 inorder 中的下标
	for {
		if in[rootIndex] == root.Val {
			break
		}
		rootIndex++
	}

	// rootIndex - inLeft 表示左子树的点数
	root.Left = build(preLeft+1, preLeft+(rootIndex-inLeft), inLeft, rootIndex-1)
	root.Right = build(preLeft+(rootIndex-inLeft)+1, preRight, rootIndex+1, inRight)

	return root
}

func TestBuildTree(t *testing.T) {

	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	t.Log(root)

}
