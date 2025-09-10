package lru_test

import (
	"testing"

	dl "github.com/andruixxd31/kata-machine-go/src/doubly_linked_list"
	"github.com/andruixxd31/kata-machine-go/src/lru"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		search string
		want   *dl.Node
		cache  *lru.LRU
	}{
		{
			name:   "empty LRU cache should return nil",
			search: "file1",
			want:   nil,
			cache: &lru.LRU{
				Capacity:   2,
				LinkedList: &dl.LinkedList{},
				Map:        map[string]*dl.Node{},
			},
		},
		{
			name:   "hit in cache returns node and moves it to tail",
			search: "node1",
			want:   &dl.Node{Key: "node1", Val: 100},
			cache: func() *lru.LRU {
				ll := &dl.LinkedList{}
				m := make(map[string]*dl.Node)

				// Create nodes
				node1 := &dl.Node{Key: "node1", Val: 100}
				node2 := &dl.Node{Key: "node2", Val: 200}
				node3 := &dl.Node{Key: "node3", Val: 300}

				ll.Append(node1)
				ll.Append(node2)
				ll.Append(node3)

				// Fill the map
				m["node1"] = node1
				m["node2"] = node2
				m["node3"] = node3

				return &lru.LRU{
					Capacity:   3,
					LinkedList: ll,
					Map:        m,
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.cache.Get(tt.search)

			assertEqual(t, got, tt.want)

			if got != nil && tt.cache.LinkedList.Tail.Key != got.Key {
				t.Errorf("%s: expected node %+v to be tail, but tail is %+v", tt.name, got.Key, tt.cache.LinkedList.Tail.Key)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name   string
		update *dl.Node
		want   *dl.Node
		cache  *lru.LRU
	}{
		{
			name:   "Update inside empty lru",
			update: &dl.Node{Key: "node1", Val: 100},
			want:   &dl.Node{Key: "node1", Val: 100},
			cache: &lru.LRU{
				Capacity:   2,
				LinkedList: &dl.LinkedList{},
				Map:        map[string]*dl.Node{},
			},
		},
		{
			name:   "add item in lru with cap maxed",
			update: &dl.Node{Key: "node4", Val: 400},
			want:   &dl.Node{Key: "node4", Val: 400},
			cache: func() *lru.LRU {
				ll := &dl.LinkedList{}
				m := make(map[string]*dl.Node)

				// Create nodes
				node1 := &dl.Node{Key: "node1", Val: 100}
				node2 := &dl.Node{Key: "node2", Val: 200}
				node3 := &dl.Node{Key: "node3", Val: 300}

				ll.Append(node1)
				ll.Append(node2)
				ll.Append(node3)

				// Fill the map
				m["node1"] = node1
				m["node2"] = node2
				m["node3"] = node3

				return &lru.LRU{
					Capacity:   3,
					LinkedList: ll,
					Map:        m,
				}
			}(),
		},
		{
			name:   "update existing node",
			update: &dl.Node{Key: "node2", Val: 400},
			want:   &dl.Node{Key: "node2", Val: 400},
			cache: func() *lru.LRU {
				ll := &dl.LinkedList{}
				m := make(map[string]*dl.Node)

				// Create nodes
				node1 := &dl.Node{Key: "node1", Val: 100}
				node2 := &dl.Node{Key: "node2", Val: 200}
				node3 := &dl.Node{Key: "node3", Val: 300}

				ll.Append(node1)
				ll.Append(node2)
				ll.Append(node3)

				// Fill the map
				m["node1"] = node1
				m["node2"] = node2
				m["node3"] = node3

				return &lru.LRU{
					Capacity:   3,
					LinkedList: ll,
					Map:        m,
				}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache.Update(tt.update.Key, tt.update)

			assertEqual(t, tt.cache.LinkedList.Tail, tt.want)

			assertCapIsNotOverFlown(t, tt.cache.LinkedList.Length, tt.cache.Capacity)

			assertEqual(t, tt.cache.Map[tt.update.Key], tt.cache.Map[tt.want.Key])

		})
	}

}

func assertEqual(t testing.TB, got, want *dl.Node) {
	t.Helper()

	if got == nil && want == nil {
		return
	}
	if got == nil || want == nil {
		t.Errorf("got %+v, want %+v", got, want)
		return
	}

	if got.Key != want.Key || got.Val != want.Val {
		t.Errorf("got node {Key: %s, Val: %d}, want node {Key: %s, Val: %d}", got.Key, got.Val, want.Key, want.Val)
	}
}

func assertCapIsNotOverFlown(t testing.TB, got, want int) {
	t.Helper()

	if got > want {
		t.Errorf("got: %d is bigger then want: %d", got, want)
		return
	}
}
