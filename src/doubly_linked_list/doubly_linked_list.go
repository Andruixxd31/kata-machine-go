package doublylinkedlist

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

type Node struct {
	Key  string
	Val  int
	Next *Node
	Prev *Node
}

func (dl *LinkedList) Append(node *Node) {
	dl.Length++
	if dl.Tail == nil {
		dl.Tail = node
		dl.Head = node
		return
	}

	dl.Tail.Next = node
	node.Prev = dl.Tail
	dl.Tail = node
}

func (dl *LinkedList) Prepend(node *Node) {
	dl.Length++
	if dl.Head == nil {
		dl.Head = node
		dl.Tail = node
		return
	}

	dl.Head.Prev = node
	node.Next = dl.Head
	dl.Head = node
}

func (dl *LinkedList) InsertAt(idx int, node *Node) {
	if idx < 0 || idx > dl.Length {
		return
	}

	if idx == dl.Length {
		dl.Append(node)
		return
	}

	if idx == 0 {
		dl.Prepend(node)
		return
	}

	count := 0
	cur := dl.Head

	for count < idx {
		cur = cur.Next
		count++
	}

	prev := cur.Prev

	node.Next = cur
	node.Prev = prev
	cur.Prev = node
	prev.Next = node

	dl.Length++
}

func (dl *LinkedList) DeleteLast() {
	if dl.Head == nil {
		return
	}

	if dl.Head == dl.Tail {
		// only one node
		dl.Head = nil
		dl.Tail = nil
		dl.Length = 0
		return
	}

	node := dl.Head
	dl.Head = node.Next
	dl.Head.Prev = nil
	node.Next = nil

	dl.Length--
}
