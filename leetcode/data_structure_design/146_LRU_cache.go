package leetcode

// https://leetcode.com/problems/lru-cache/

// Use a linked list combined with a map
// Map for instant access
// Linked list means a node can move to top instantly
// and can remove from 1 of the side instantly
// node contains [prev_node, key, next_node]
// Map [key] to [value, node]
// When put
// if size == capacity
// pop the last node and delete the key
// then add to map[key] and add node to beginning
// When get
// Since we have reference to the node, we can instantly move it to the beginning
// previous = node.prev_node
// next = node.next_node
// previous.next_node = next
// next.prev_node = previous
// node.prev_node = null, node.next_node = first_node
//
// Use an extra 'start' and 'end' node so we don't need any nil checks
type Node struct {
	next  *Node
	prev  *Node
	key   int
	value int
}
type LRUCache struct {
	keys  map[int]*Node
	start *Node
	end   *Node
	cap   int
	count int
}

func Constructor(capacity int) LRUCache {
	start := &Node{}
	end := &Node{}
	start.next = end
	end.prev = start
	return LRUCache{
		keys:  make(map[int]*Node),
		cap:   capacity,
		end:   end,
		start: start,
	}
}

func (cache *LRUCache) Get(key int) int {
	node := cache.keys[key]
	if node == nil {
		return -1
	}
	// pull the node out and link the other ones together
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev

	// bring it to the start, have an extra 'start' node to avoid nil checks
	nextForThis := cache.start.next
	node.prev = cache.start
	node.next = nextForThis
	nextForThis.prev = node
	cache.start.next = node

	return node.value
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.keys[key] != nil {
		cache.keys[key].value = value
		cache.Get(key)
		return
	}
	if cache.count == cache.cap {
		evictNode := cache.end.prev
		evictNode.prev.next = cache.end
		cache.end.prev = evictNode.prev
		cache.keys[evictNode.key] = nil
	} else {
		cache.count++
	}
	nextForThis := cache.start.next
	node := &Node{
		key:   key,
		value: value,
		prev:  cache.start,
		next:  nextForThis,
	}
	nextForThis.prev = node
	cache.start.next = node
	cache.keys[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
