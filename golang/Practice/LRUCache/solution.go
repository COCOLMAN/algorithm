package LRUCache

type nodeItem struct {
	key    int
	value  int
	next   *nodeItem
	before *nodeItem
}

type LRUCache struct {
	head            *nodeItem
	tail            *nodeItem
	values          map[int]*nodeItem
	capacity        int
	currentCapacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:            nil,
		tail:            nil,
		capacity:        capacity,
		currentCapacity: 0,
		values:          map[int]*nodeItem{},
	}
}

func (this *LRUCache) makeRecently(key int) {
	targetNode := this.values[key]
	if this.head.key == key {
		return
	}
	if this.tail.key == key {
		targetNode.before.next = nil
		this.tail = targetNode.before
		targetNode.before = nil
		this.head.before = targetNode
		targetNode.next = this.head
		this.head = targetNode
		return
	}

	targetNode.before.next = targetNode.next
	targetNode.next.before = targetNode.before
	this.head.before = targetNode
	targetNode.next = this.head
	this.head = targetNode
}

func (this *LRUCache) Get(key int) int {
	val, exist := this.values[key]
	if !exist {
		return -1
	}
	this.makeRecently(key)
	return val.value
}

func (this *LRUCache) Put(key int, value int) {
	// if exists key, update value and re-order
	_, exist := this.values[key]
	if exist {
		this.values[key].value = value
		this.makeRecently(key)
		return
	}

	newNode := &nodeItem{
		key:   key,
		value: value,
	}
	this.values[key] = newNode
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
