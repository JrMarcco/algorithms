package deleteNode_450

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var toDeleteParent *TreeNode
	toDeleteNode := root // 值匹配的节点，即待删除的节点
	for toDeleteNode != nil && toDeleteNode.Val != key {
		toDeleteParent = toDeleteNode
		if key < toDeleteNode.Val {
			toDeleteNode = toDeleteNode.Left
		} else {
			toDeleteNode = toDeleteNode.Right
		}
	}

	if toDeleteNode == nil {
		// 未找到值匹配的节点
		return root
	}

	if toDeleteNode.Left == nil && toDeleteNode.Right == nil {
		toDeleteNode = nil
	} else if toDeleteNode.Left == nil {
		toDeleteNode = toDeleteNode.Right
	} else if toDeleteNode.Right == nil {
		toDeleteNode = toDeleteNode.Left
	} else {
		// 寻找后继节点
		successorParent := toDeleteNode // 后继节点的父节点
		successor := toDeleteNode.Right // 后继节点
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}

		if successorParent.Val == toDeleteNode.Val {
			// 后继节点没有左子树的情况
			successorParent.Right = successor.Right
		} else {
			successorParent.Left = successor.Right
		}

		// 用后继节点替换待删除节点
		successor.Left = toDeleteNode.Left
		successor.Right = toDeleteNode.Right
		toDeleteNode = successor
	}

	if toDeleteParent == nil {
		return toDeleteNode
	}
	if toDeleteParent.Left != nil && toDeleteParent.Left.Val == key {
		toDeleteParent.Left = toDeleteNode
	} else {
		toDeleteParent.Right = toDeleteNode
	}
	return root
}

func deleteNode_(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 查找后继节点，
		successor := root.Right
		for ; successor.Left != nil; successor = successor.Left {
		}
		// 将后继节点从右子树中删除
		root.Right = deleteNode(root.Right, successor.Val)
		// 用后继节点的值替换根节点的值
		root.Val = successor.Val
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func TestDeleteNode(t *testing.T) {

	root := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
	}

	t.Log(deleteNode(root, 2))
}
