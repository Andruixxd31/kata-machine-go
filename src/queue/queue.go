package queue

import (
	"fmt"
	"strings"
)

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Val  int
	Next *Node
}

func (q *Queue) Enqueue(node *Node) {
	q.Length++

	if q.Head == nil {
		q.Head = node
		q.Tail = node
		return
	}

	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue) Dequeue() Node {
	if q.Head == nil {
		return Node{}
	}

	q.Length--

	if q.Length == 0 {
		node := q.Head
		q.Head = nil
		q.Tail = nil
		return *node
	}

	node := q.Head
	q.Head = q.Head.Next
	node.Next = nil

	return *node
}

func (q *Queue) Peek() *Node {
	return q.Head
}

func (q Queue) PrintQueue() {
	var qString strings.Builder
	cur := q.Head

	for cur != nil {
		qString.WriteString(fmt.Sprintf("%d", cur.Val))
		if cur.Next != nil {
			qString.WriteString(" -> ")
		} else {
			fmt.Fprintln(&qString)
		}
		cur = cur.Next
	}

	fmt.Print(qString.String())
}
