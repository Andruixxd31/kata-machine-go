package lru

import (
	dl "github.com/andruixxd31/kata-machine-go/src/doubly_linked_list"
)

type LRU struct {
	Capacity   int
	LinkedList *dl.LinkedList
	Map        map[string]*dl.Node
}

func (lru *LRU) Update(key string, value *dl.Node) {
	if node, ok := lru.Map[key]; ok {
		// ✅ Happy path: key exists
		node.Val = value.Val
		lru.LinkedList.DeleteNode(node)
		lru.LinkedList.Append(node)
		return
	}

	// ❌ Key doesn't exist: create new node
	node := &dl.Node{Key: key, Val: value.Val}
	lru.LinkedList.Append(node)
	lru.Map[key] = node

	// Evict oldest if over capacity
	if lru.LinkedList.Length > lru.Capacity {
		oldest := lru.LinkedList.Head
		lru.LinkedList.DeleteNode(oldest)
		delete(lru.Map, oldest.Key)
	}
}

// Returns the node if it exists.
// Moves node to latest used position in linked list
func (lru *LRU) Get(key string) *dl.Node {
	//check if there is a registry of node
	value, ok := lru.Map[key]
	if !ok {
		return nil
	}

	// Move node to tail
	lru.LinkedList.DeleteNode(value)
	lru.LinkedList.Append(value)

	return value
}
