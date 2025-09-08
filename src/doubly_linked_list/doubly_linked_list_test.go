package doublylinkedlist_test

import (
	"reflect"
	"testing"

	doublylinkedlist "github.com/andruixxd31/kata-machine-go/src/doubly_linked_list"
)

func TestDoublyLinkedListAppend(t *testing.T) {
	tests := []struct {
		name       string
		ll         *doublylinkedlist.LinkedList
		want       *doublylinkedlist.LinkedList
		valueToAdd int
	}{
		{
			name: "Append to empty doubly linked list",
			ll:   &doublylinkedlist.LinkedList{},
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			valueToAdd: 1,
		},
		{
			name: "Append to non-empty doubly linked list",
			ll: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node2 := &doublylinkedlist.Node{Val: 2}
				node1.Next = node2
				node2.Prev = node1
				return &doublylinkedlist.LinkedList{
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
			tt.ll.Append(&doublylinkedlist.Node{Val: tt.valueToAdd})

			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("got %+v, want %+v", tt.ll, tt.want)
			}
		})
	}
}

func TestDoublyLinkedListPrepend(t *testing.T) {
	tests := []struct {
		name           string
		ll             *doublylinkedlist.LinkedList
		want           *doublylinkedlist.LinkedList
		valueToPrepend int
	}{
		{
			name: "Prepend to empty list",
			ll:   &doublylinkedlist.LinkedList{},
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			valueToPrepend: 1,
		},
		{
			name: "Prepend to non-empty list",
			ll: func() *doublylinkedlist.LinkedList {
				node2 := &doublylinkedlist.Node{Val: 2}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node2,
					Tail:   node2,
				}
			}(),
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node2 := &doublylinkedlist.Node{Val: 2}
				node1.Next = node2
				node2.Prev = node1
				return &doublylinkedlist.LinkedList{
					Length: 2,
					Head:   node1,
					Tail:   node2,
				}
			}(),
			valueToPrepend: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.Prepend(&doublylinkedlist.Node{Val: tt.valueToPrepend})

			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("got %+v, want %+v", tt.ll, tt.want)
			}
		})
	}
}

func TestDoublyLinkedListInsertAt(t *testing.T) {
	tests := []struct {
		name  string
		ll    *doublylinkedlist.LinkedList
		want  *doublylinkedlist.LinkedList
		idx   int
		value int
	}{
		{
			name: "Insert at head (index 0)",
			ll: func() *doublylinkedlist.LinkedList {
				node2 := &doublylinkedlist.Node{Val: 2}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node2,
					Tail:   node2,
				}
			}(),
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node2 := &doublylinkedlist.Node{Val: 2}
				node1.Next = node2
				node2.Prev = node1
				return &doublylinkedlist.LinkedList{
					Length: 2,
					Head:   node1,
					Tail:   node2,
				}
			}(),
			idx:   0,
			value: 1,
		},
		{
			name: "Insert in middle",
			ll: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node3 := &doublylinkedlist.Node{Val: 3}
				node1.Next = node3
				node3.Prev = node1
				return &doublylinkedlist.LinkedList{
					Length: 2,
					Head:   node1,
					Tail:   node3,
				}
			}(),
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node2 := &doublylinkedlist.Node{Val: 2}
				node3 := &doublylinkedlist.Node{Val: 3}
				node1.Next = node2
				node2.Prev = node1
				node2.Next = node3
				node3.Prev = node2
				return &doublylinkedlist.LinkedList{
					Length: 3,
					Head:   node1,
					Tail:   node3,
				}
			}(),
			idx:   1,
			value: 2,
		},
		{
			name: "Insert at tail (index == length)",
			ll: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			want: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node2 := &doublylinkedlist.Node{Val: 2}
				node1.Next = node2
				node2.Prev = node1
				return &doublylinkedlist.LinkedList{
					Length: 2,
					Head:   node1,
					Tail:   node2,
				}
			}(),
			idx:   1,
			value: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.InsertAt(tt.idx, &doublylinkedlist.Node{Val: tt.value})

			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("got %+v, want %+v", tt.ll, tt.want)
			}
		})
	}
}

func TestDoublyLinkedListDeleteLast(t *testing.T) {
	tests := []struct {
		name string
		ll   *doublylinkedlist.LinkedList
		want *doublylinkedlist.LinkedList
	}{
		{
			name: "Delete from empty list",
			ll:   &doublylinkedlist.LinkedList{},
			want: &doublylinkedlist.LinkedList{},
		},
		{
			name: "Delete from single-node list",
			ll: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				return &doublylinkedlist.LinkedList{
					Length: 1,
					Head:   node1,
					Tail:   node1,
				}
			}(),
			want: &doublylinkedlist.LinkedList{},
		},
		{
			name: "Delete from multi-node list",
			ll: func() *doublylinkedlist.LinkedList {
				node1 := &doublylinkedlist.Node{Val: 1}
				node2 := &doublylinkedlist.Node{Val: 2}
				node3 := &doublylinkedlist.Node{Val: 3}
				node1.Next = node2
				node2.Prev = node1
				node2.Next = node3
				node3.Prev = node2
				return &doublylinkedlist.LinkedList{
					Length: 3,
					Head:   node1,
					Tail:   node3,
				}
			}(),
			want: func() *doublylinkedlist.LinkedList {
				node2 := &doublylinkedlist.Node{Val: 2}
				node3 := &doublylinkedlist.Node{Val: 3}
				node2.Next = node3
				node3.Prev = node2
				return &doublylinkedlist.LinkedList{
					Length: 2,
					Head:   node2,
					Tail:   node3,
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.DeleteLast()

			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("got %+v, want %+v", tt.ll, tt.want)
			}
		})
	}
}
