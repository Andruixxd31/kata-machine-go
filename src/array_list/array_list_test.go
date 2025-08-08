package arraylist_test

import (
	"fmt"
	"reflect"
	"testing"

	arraylist "github.com/andruixxd31/kata-machine-go/src/array_list"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		name string
		cap  int
		len  int
		got  []int
		want []int
	}{
		{
			name: "append within len",
			cap:  2,
			len:  1,
			got:  make([]int, 0, 2),
			want: []int{1},
		},
		{
			name: "append maxing cap",
			cap:  1,
			len:  1,
			got:  make([]int, 0, 0),
			want: []int{1},
		},
		{
			name: "append maxing cap 2",
			cap:  7,
			len:  4,
			got:  []int{1, 2, 3},
			want: []int{1, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayList := arraylist.Append(tt.got, 1)

			if len(arrayList) != tt.len {
				t.Errorf("len is different: got %d - want %d", len(arrayList), tt.len)
			}

			if cap(arrayList) != tt.cap {
				t.Errorf("cap is different: got %d - want %d", cap(arrayList), tt.cap)
			}

			if !reflect.DeepEqual(arrayList, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.got, tt.got, tt.want)
			}

			fmt.Printf("%#v \n", arrayList)
		})
	}
}
