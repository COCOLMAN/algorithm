package LRUCache

import "fmt"

type nodeItem struct {
	key    int
	val    int
	next   *nodeItem
	before *nodeItem
}

type LRUCache struct {
	head        *nodeItem
	capacity    int
	currentSize int
}

func printData(cache LRUCache) {
	result := ""
	node := cache.head.next
	for node != nil {
		result += fmt.Sprintf("%d:%d ", node.key, node.val)
		node = node.next
	}
	fmt.Println(result, cache.currentSize)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:    capacity,
		currentSize: 0,
		head:        &nodeItem{},
	}
}

func link(front, back *nodeItem) {
	if back == nil {
		return
	}
	front.next = back
	back.before = front
}

func (this *LRUCache) Get(key int) int {
	node := this.head
	var targetNode *nodeItem
	for node.next != nil {
		if node.next.key == key {
			targetNode = node.next
			link(targetNode.before, targetNode.next)
		}
		node = node.next
	}
	if targetNode == nil {
		return -1
	}
	node.next = targetNode
	targetNode.before = node
	targetNode.next = nil

	return targetNode.val
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == this.currentSize {
		n := this.head.next

		this.head.next = n.next
		if n.next != nil {
			n.next.before = this.head
		}
		this.currentSize -= 1
	}

	if this.head.next == nil {
		this.currentSize += 1
		this.head.next = &nodeItem{
			key:    key,
			val:    value,
			next:   nil,
			before: this.head,
		}
		return
	}

	node := this.head
	for node.next != nil {
		if node.key == key {
			node.val = value
			return
		}
		node = node.next
	}
	this.currentSize += 1
	node.next = &nodeItem{
		key:    key,
		val:    value,
		before: node,
	}
}
