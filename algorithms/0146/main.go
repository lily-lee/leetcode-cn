package _0146

type LRUCache struct {
	cap  int
	m    map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		cap:  capacity,
		m:    map[int]*Node{},
		head: &Node{},
		tail: &Node{},
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}
	this.addToFirst(node)
	return node.value

}
func (this *LRUCache) addToFirst(node *Node) {
	if this.head.next == node {
		return
	}
	firstOld := this.head.next
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	firstOld.prev = node
	node.next = firstOld
	node.prev = this.head
	this.head.next = node
	this.m[node.key] = node
}

func (this *LRUCache) removeLast() {
	last := this.tail.prev
	delete(this.m, last.key)
	prev := last.prev
	this.tail.prev = prev
	prev.next = this.tail
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.value = value
	} else {
		node = &Node{key: key, value: value}
	}
	this.addToFirst(node)
	this.m[key] = node

	if len(this.m) > this.cap {
		this.removeLast()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
