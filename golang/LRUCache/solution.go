package LRUCache

import "fmt"

func printCache(c LRUCache) {
	fmt.Println("--------- start ")
	fmt.Printf("%d / %d\n", c.currentCapacity, c.capacity)
	fmt.Println(c.values)
	result := ""
	node := c.head
	for node != nil {
		result += fmt.Sprintf(" %d:%d", node.key, c.values[node.key])
		node = node.next
	}
	fmt.Println(result)
	fmt.Println("---------")
}

type nodeItem struct {
	key    int
	next   *nodeItem
	before *nodeItem
}

type LRUCache struct {
	head            *nodeItem
	tail            *nodeItem
	values          map[int]int
	capacity        int
	currentCapacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:            nil,
		tail:            nil,
		capacity:        capacity,
		currentCapacity: 0,
		values:          map[int]int{},
	}
}

func (this *LRUCache) makeRecently(key int) {
	if this.head.key == this.tail.key {
		return
	}
	if this.tail.key == key {
		beforeTail := this.tail
		if beforeTail.before != nil {
			beforeTail.before.next = nil
		}
		this.tail = beforeTail.before

		beforeHead := this.head
		this.head = beforeTail
		beforeTail.before = nil
		beforeTail.next = beforeHead
		beforeHead.before = beforeTail
		return
	}
	if this.head.key == key {
		return
	}

	// this.tail이 바뀌지 않음
	var targetNode *nodeItem
	node := this.head
	for node != nil {
		if node.key == key {
			targetNode = node
			if targetNode.before != nil {
				targetNode.before.next = targetNode.next
			}
			if targetNode.next != nil {
				targetNode.next.before = targetNode.before
			}
		}
		node = node.next
	}
	beforeHead := this.head
	this.head = targetNode
	targetNode.before = nil
	targetNode.next = beforeHead
	beforeHead.before = targetNode
}

func (this *LRUCache) Get(key int) int {
	val, exist := this.values[key]
	if !exist {
		return -1
	}
	this.makeRecently(key)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	// if exists key, update value and re-order
	_, exist := this.values[key]
	if exist {
		this.values[key] = value
		this.makeRecently(key)
		return
	}

	newNode := &nodeItem{
		key: key,
	}
	this.values[key] = value
	this.currentCapacity++
	// if empty node, set first node and set head, tail
	if this.currentCapacity == 1 {
		this.head = newNode
		this.tail = newNode
		return
	}
	// if not exist key, set value and set node
	beforeHead := this.head
	newNode.next = beforeHead
	beforeHead.before = newNode
	this.head = newNode
	// if over capacity
	if this.capacity < this.currentCapacity {
		beforeTail := this.tail
		this.tail = beforeTail.before
		this.tail.next = nil
		this.currentCapacity--
		delete(this.values, beforeTail.key)
	}
}
