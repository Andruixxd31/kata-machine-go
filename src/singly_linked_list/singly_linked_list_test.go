package singlylinkedlist_test

import (
	"testing"

	singlylinkedlist "github.com/andruixxd31/kata-machine-go/src/singly_linked_list"
)

// TODO: Add a benchark on print linked list and see if other case is aplicable
func BenchmarkSinglyLinkedListAppend(b *testing.B) {

	var ll singlylinkedlist.LinkedList
	var node singlylinkedlist.Node
	node.Val = 1
	node4 := singlylinkedlist.Node{
		Val: 4,
	}
	node2 := singlylinkedlist.Node{
		Val: 2,
	}

	ll.PrintLinkedList()
	ll.Append(&node)
	ll.Append(&node2)
	ll.Append(&node4)

	for b.Loop() {
		ll.PrintLinkedList()
	}

}

// func TestSinglyLinkedListAppend(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		ll   singlylinkedlist.LinkedList
// 		want singlylinkedlist.LinkedList
// 	}{
// 		{
// 			name: "Append a single element",
// 			ll:   singlylinkedlist.LinkedList{},
// 			want: singlylinkedlist.LinkedList{
// 				Head: &singlylinkedlist.Node{
// 					Val:  1,
// 					Next: nil,
// 				},
// 				Tail: &singlylinkedlist.Node{
// 					Val:  1,
// 					Next: nil,
// 				},
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
//
// 	}
// }
