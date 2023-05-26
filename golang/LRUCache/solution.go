package LRUCache

type nodeItem struct {
	key  int
	val  int
	next *nodeItem
}

type LRUCache struct {
	node     *nodeItem
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.node
	//var beforeNode *nodeItem
	for node != nil {
		if node.key == key {
			return node.val
			//if beforeNode != nil {
			//	beforeNode.next = node.next
			//}
		}
		//beforeNode = node
		node = node.next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.node == nil {
		this.node = &nodeItem{
			key:  key,
			val:  value,
			next: nil,
		}
		return
	}

	node := this.node
	for node.next != nil {
		if node.key == key {
			node.val = value
			return
		}
		node = node.next
	}
	node.next = &nodeItem{
		key: key,
		val: value,
	}
}
