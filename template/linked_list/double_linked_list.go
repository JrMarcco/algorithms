package linked_list

type LinkedListNode struct {
	Value    int
	PreNode  *LinkedListNode
	NextNode *LinkedListNode
}

// InsertAfter 在某一节点后插入新节点
func InsertAfter(node *LinkedListNode, value int) *LinkedListNode {
	newNode := &LinkedListNode{Value: value}

	node.NextNode.PreNode = newNode
	newNode.NextNode = node.NextNode

	node.NextNode = newNode
	newNode.PreNode = node

	return newNode
}

// InsertBefore 在某一节点前插入新节点
func InsertBefore(node *LinkedListNode, value int) *LinkedListNode {
	newNode := &LinkedListNode{Value: value}

	node.PreNode.NextNode = newNode
	newNode.PreNode = node.PreNode

	node.PreNode = newNode
	newNode.NextNode = node

	return newNode
}
