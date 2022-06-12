package lru

// Cache 使用链表存储数据，map 记录 key 对应的链表节点
type Cache struct {
	cacheNodeMap map[int]*cacheNode
	cap          int
	head         *cacheNode
	tail         *cacheNode
}

// 缓存链表节点
type cacheNode struct {
	key  int
	val  int
	prev *cacheNode
	next *cacheNode
}

func New(capacity int) Cache {

	head := &cacheNode{}
	tail := &cacheNode{}

	head.next = tail
	tail.prev = head

	return Cache{
		cacheNodeMap: map[int]*cacheNode{},
		cap:          capacity,
		head:         head,
		tail:         tail,
	}
}

func (lru *Cache) Get(key int) int {
	if node, ok := lru.cacheNodeMap[key]; ok {
		node.prev.next = node.next
		node.next.prev = node.prev

		// 将元素放入队尾
		lru.mvToTail(node)

		return node.val
	}
	return -1
}

func (lru *Cache) Put(key int, value int) {
	if node, ok := lru.cacheNodeMap[key]; ok {
		node.val = value
		lru.removeCacheNode(node)
		lru.mvToTail(node)
		return
	}

	if len(lru.cacheNodeMap) == lru.cap {
		// 移除头元素
		lru.removeHead()
	}

	newNode := &cacheNode{
		key: key,
		val: value,
	}

	lru.mvToTail(newNode)
	lru.cacheNodeMap[key] = newNode
}

func (lru *Cache) removeHead() {
	delete(lru.cacheNodeMap, lru.head.next.key)

	lru.removeCacheNode(lru.head.next)
}

func (lru *Cache) removeCacheNode(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *Cache) mvToTail(node *cacheNode) {
	// 当前元素移到队尾
	node.prev = lru.tail.prev
	node.next = lru.tail
	lru.tail.prev.next = node
	lru.tail.prev = node
}
