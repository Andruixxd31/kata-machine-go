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
			tt.bst.Insert(tt.node)
			if !reflect.DeepEqual(tt.bst, tt.want) {
				spew.Dump(tt.bst)
				spew.Dump(tt.want)
				t.Errorf("Want %+v - got %+v", &tt.bst, &tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
		val  int      // value to delete
		bst  *bst.BST // input tree
		want *bst.BST // expected tree after deletion
	}{
		{
			name: "Delete leaf node",
			val:  0,
			bst: &bst.BST{
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
			want: &bst.BST{
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
		},
		{
			name: "Delete node with one child",
			val:  1,
			bst: &bst.BST{
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
			want: &bst.BST{
				Root: &bst.Node{
					Val: 5,
					Left: &bst.Node{
						Val: 0,
					},
					Right: &bst.Node{
						Val: 10,
					},
				},
			},
		},
		{
			name: "Delete node with two children",
			val:  5,
			bst: &bst.BST{
				Root: &bst.Node{
					Val: 5,
					Left: &bst.Node{
						Val: 3,
					},
					Right: &bst.Node{
						Val: 7,
						Left: &bst.Node{
							Val: 6,
						},
						Right: &bst.Node{
							Val: 9,
						},
					},
				},
			},
			want: &bst.BST{
				Root: &bst.Node{
					Val: 6, // successor of 5
					Left: &bst.Node{
						Val: 3,
					},
					Right: &bst.Node{
						Val: 7,
						Right: &bst.Node{
							Val: 9,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bst.Delete(&bst.Node{Val: tt.val})

			// Deep equality check
			if !reflect.DeepEqual(tt.bst, tt.want) {
				t.Errorf("Delete(%d) failed:\n got  %+v\n want %+v",
					tt.val, tt.bst, tt.want)
			}
		})
	}
}
