package singlylinkedlist_test

import (
	"testing"

	singlylinkedlist "github.com/andruixxd31/kata-machine-go/src/singly_linke_list"
)

func TestSinglyLinkedListAppend(t *testing.T) {
	tests := []struct {
		name string
		ll   singlylinkedlist.LinkedList
		want singlylinkedlist.LinkedList
	}{
		{
			name: "Append a single element",
			ll:   singlylinkedlist.LinkedList{},
			want: singlylinkedlist.LinkedList{
				Head: &singlylinkedlist.Node{
					Val:  1,
					Next: nil,
				},
				Tail: &singlylinkedlist.Node{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	for _, tt := range tests {

	}
}
