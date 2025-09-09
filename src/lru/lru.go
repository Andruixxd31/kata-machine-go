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
