package singlylinkedlist_test

import (
	"reflect"
	"testing"

	singlylinkedlist "github.com/andruixxd31/kata-machine-go/src/singly_linked_list"
)

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

func TestSinglyLinkedListAppend(t *testing.T) {
	tests := []struct {
		name       string
		ll         *singlylinkedlist.LinkedList
		want       *singlylinkedlist.LinkedList
		valueToAdd int
	}{
		{
			name: "Append to empty singly linked list",
			ll:   &singlylinkedlist.LinkedList{},
			want: func() *singlylinkedlist.LinkedList {
				node1 := &singlylinkedlist.Node{Val: 1}
				return &singlylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			valueToAdd: 1,
		},
		{
			name: "Append to singly linked list with elements",
			ll: func() *singlylinkedlist.LinkedList {
				node1 := &singlylinkedlist.Node{Val: 1}
				return &singlylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			want: func() *singlylinkedlist.LinkedList {
				node1 := &singlylinkedlist.Node{Val: 1}
				node2 := &singlylinkedlist.Node{Val: 2}
				node1.Next = node2
				return &singlylinkedlist.LinkedList{
					Length: 2,
					Head:   node1,
					Tail:   node2,
				}
			}(),
			valueToAdd: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.Append(&singlylinkedlist.Node{Val: tt.valueToAdd})

			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("got %+v, want %+v", tt.ll, tt.want)
			}
		})
	}
}
