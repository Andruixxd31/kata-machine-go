package bst_test

import (
	"testing"

	"github.com/andruixxd31/kata-machine-go/src/bst"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name string
		node bst.Node
		bst  *bst.BST
		want bool
	}{
		{
			name: "Node is present",
			node: bst.Node{
				Val: 1,
			},
			bst: &bst.BST{
				Root: &bst.Node{
					Val: 5,
					Left: &bst.Node{
						Val: 1,
					},
					Right: &bst.Node{
						Val: 10,
					},
				},
			},
			want: true,
		},
		{
			name: "Node is not present",
			node: bst.Node{
				Val: 1,
			},
			bst: &bst.BST{
				Root: &bst.Node{
					Val: 5,
					Left: &bst.Node{
						Val: 2,
					},
					Right: &bst.Node{
						Val: 10,
					},
				},
			},
			want: false,
		},
		{
			name: "Empty bst",
			node: bst.Node{
				Val: 1,
			},
			bst:  &bst.BST{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bst.Find(tt.node)
			if got != tt.want {
				t.Errorf("Want %t - got %t", got, tt.want)
			}
		})
	}
}
