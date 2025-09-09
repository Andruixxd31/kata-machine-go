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

func (lru *LRU) Get(key string) *dl.Node {
	// check cache for existance
	var item dl.Node
	lru.Map["node1"] = &dl.Node{Key: "node1", Val: 100}

	value, ok := lru.Map[key]
	if !ok {
		return nil
	} else {
		item = *value
	}

	// update value and move it to the front

	// return value of nil
	return &item
}
