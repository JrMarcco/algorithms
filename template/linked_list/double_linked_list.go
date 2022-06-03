package linked_list

type LinkedListNode struct {
	Value    int
	PrevNode *LinkedListNode
	NextNode *LinkedListNode
}

// InsertAfter 在某一节点后插入新节点
func InsertAfter(node *LinkedListNode, value int) *LinkedListNode {
	newNode := &LinkedListNode{Value: value}

	node.NextNode.PrevNode = newNode
	newNode.NextNode = node.NextNode

	node.NextNode = newNode
	newNode.PrevNode = node

	return newNode
}

// InsertBefore 在某一节点前插入新节点
func InsertBefore(node *LinkedListNode, value int) *LinkedListNode {
	newNode := &LinkedListNode{Value: value}

	node.PrevNode.NextNode = newNode
	newNode.PrevNode = node.PrevNode

	node.PrevNode = newNode
	newNode.NextNode = node

	return newNode
}

func Delete(node *LinkedListNode) {
	node.PrevNode.NextNode = node.NextNode
	node.NextNode.PrevNode = node.PrevNode
}
