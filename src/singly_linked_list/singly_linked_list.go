package singlylinkedlist

import (
	"fmt"
	"reflect"
	"strings"
)

type SinglyLinkedList interface {
	Prepend(node *Node)
	InsertAt(node *Node, idx int)
	Append(node *Node)
	Remove(node *Node) *Node
	Get(idx int) *Node
	RemoveAt(idx int) *Node
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

type Node struct {
	Val  int
	Next *Node // has to point to next node
}

func (ll *LinkedList) Append(node *Node) {
	ll.Length++
	if ll.Tail == nil {
		ll.Tail = node
		ll.Head = node
		return
	}

	ll.Tail.Next = node
	ll.Tail = node
}

func (ll *LinkedList) Prepend(node *Node) {
	ll.Length++
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	node.Next = ll.Head
	ll.Head = node
}

func (ll *LinkedList) InsertAt(idx int, node *Node) {
	if idx < 0 {
		return
	}

	// idx is out of bounds
	if idx > ll.Length {
		return
	}

	//At Head
	if idx == 0 {
		ll.Prepend(node)
		return
	}

	//At tail
	if idx == ll.Length {
		ll.Append(node)
		return
	}

	//iterate to index
	count := 0
	cur := ll.Head
	var prev *Node
	for count < idx && cur != nil {
		prev = cur
		cur = cur.Next
		count++
	}

	//Insert somewhere in the middle
	prev.Next = node
	node.Next = cur
	ll.Length++
}

func (ll LinkedList) PrintLinkedList() {
	var llString strings.Builder
	cur := ll.Head

	for cur != nil {
		llString.WriteString(fmt.Sprintf("%d", cur.Val))
		if cur.Next != nil {
			llString.WriteString(" -> ")
		} else {
			fmt.Fprintln(&llString)
		}
		cur = cur.Next
	}

	fmt.Print(llString.String())
}

func (ll *LinkedList) Get(idx int) *Node {
	if idx < 0 {
		return &Node{}
	}

	if idx > ll.Length-1 {
		return nil
	}

	cur := ll.Head
	for i := 0; i < idx && cur != nil; i++ {
		cur = cur.Next
	}

	return cur
}

func (ll *LinkedList) Remove(node *Node) *Node {
	if ll.Head == nil {
		return nil
	}

	//Head is removed
	if reflect.DeepEqual(ll.Head, node) {
		removedNode := ll.Head
		ll.Head = ll.Head.Next
		removedNode.Next = nil
		ll.Length--

		if ll.Head == nil {
			ll.Tail = nil
		}

		return removedNode
	}

	prev := ll.Head
	cur := ll.Head.Next

	for cur != nil {
		if reflect.DeepEqual(node, cur) {
			prev.Next = cur.Next
			cur.Next = nil

			// Tail is being removed
			if cur == ll.Tail {
				ll.Tail = prev
			}

			ll.Length--
			return cur
		}
		prev = cur
		cur = cur.Next
	}

	return nil
}

func (ll *LinkedList) RemoveAt(idx int) *Node {
	if ll.Head == nil {
		return nil
	}

	if idx < 0 {
		return nil
	}

	if idx > ll.Length-1 {
		return nil
	}

	//Head is removed
	if idx == 0 {
		removedNode := ll.Head
		ll.Head = ll.Head.Next
		removedNode.Next = nil
		ll.Length--

		if ll.Head == nil {
			ll.Tail = nil
		}

		return removedNode
	}

	cur := ll.Head
	count := 0
	for count < idx-1 {
		if cur == nil || cur.Next == nil {
			return nil
		}
		cur = cur.Next
		count++
	}

	removedNode := cur.Next
	if removedNode == nil {
		return nil
	}

	cur.Next = removedNode.Next
	removedNode.Next = nil

	if removedNode == ll.Tail {
		ll.Tail = cur
	}

	ll.Length--
	return removedNode
}
