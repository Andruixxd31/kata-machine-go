package lru_test

import (
	"reflect"
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

				// Append nodes in order: node1 -> node2 -> node3
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

			assertDeepEqual(t, got, tt.want)

			if got != nil && !reflect.DeepEqual(tt.cache.LinkedList.Tail, got) {
				t.Errorf("%s: expected node %+v to be tail, but tail is %+v", tt.name, got, tt.cache.LinkedList.Tail)
			}
		})
	}
}

func assertDeepEqual(t testing.TB, got, want *dl.Node) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, false)
	}
}
