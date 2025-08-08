package main

import (
	"fmt"

	"github.com/andruixxd31/kata-machine-go/src/queue"
)

func main() {
	var q queue.Queue
	var node queue.Node
	node.Val = 1
	node4 := queue.Node{
		Val: 4,
	}
	// node2 := queue.Node{
	// 	Val: 2,
	// }
	// node0 := queue.Node{
	// 	Val: 0,
	// }

	q.Enqueue(&node)
	q.Enqueue(&node4)
	q.PrintQueue()
	returnedNode := q.Dequeue()
	q.PrintQueue()
	fmt.Println(returnedNode)
}
