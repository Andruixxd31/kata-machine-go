package heap_test

import (
	"reflect"
	"testing"

	"github.com/andruixxd31/kata-machine-go/src/heap"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		item int
		heap heap.Heap
		want heap.Heap
	}{
		{
			name: "insert into heap",
			item: 1,
			heap: []int{10, 20},
			want: []int{1, 20, 10},
		},
		{
			name: "insert into heap two",
			item: 3,
			heap: []int{2, 5, 7, 10, 12, 15},
			want: []int{2, 5, 3, 10, 12, 15, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap = tt.heap.Insert(tt.item)
			if !reflect.DeepEqual(tt.heap, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.heap, tt.heap, tt.want)
			}
		})

	}
}
