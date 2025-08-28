package bst_test

import (
	"reflect"
	"testing"

	"github.com/andruixxd31/kata-machine-go/src/bst"
	"github.com/davecgh/go-spew/spew"
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

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		node *bst.Node
		bst  *bst.BST
		want *bst.BST
	}{
		{
			name: "Insert into bst",
			node: &bst.Node{
				Val: 0,
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
			want: &bst.BST{
				Root: &bst.Node{
					Val: 5,
					Left: &bst.Node{
						Val: 1,
						Left: &bst.Node{
							Val: 0,
						},
					},
					Right: &bst.Node{
						Val: 10,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.bst, tt.want) {
				t.Errorf("Want %+v - got %+v", &tt.bst, &tt.want)
			}
		})
	}
}
