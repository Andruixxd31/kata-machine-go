package arraylist_test

import (
	"reflect"
	"testing"

	arraylist "github.com/andruixxd31/kata-machine-go/src/array_list"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		name       string
		cap        int
		len        int
		addElement any
		got        []any
		want       []any
	}{
		{
			name:       "append within len",
			cap:        2,
			len:        1,
			addElement: 1,
			got:        make([]any, 0, 2),
			want:       []any{1},
		},
		{
			name:       "append maxing cap",
			cap:        1,
			len:        1,
			addElement: 1,
			got:        make([]any, 0, 0),
			want:       []any{1},
		},
		{
			name:       "append to string slice",
			cap:        7,
			len:        4,
			addElement: "cash",
			got:        []any{"andres", "rufino", "bingo"},
			want:       []any{"andres", "rufino", "bingo", "cash"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayList := arraylist.Append(tt.got, tt.addElement)

			if len(arrayList) != tt.len {
				t.Errorf("len is different: got %d - want %d", len(arrayList), tt.len)
			}

			if cap(arrayList) != tt.cap {
				t.Errorf("cap is different: got %d - want %d", cap(arrayList), tt.cap)
			}

			if !reflect.DeepEqual(arrayList, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.got, tt.got, tt.want)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name string
		cap  int
		len  int
		got  []any
		want []any
	}{
		{
			name: "prepend empty",
			cap:  1,
			len:  1,
			got:  make([]any, 0, 0),
			want: []any{1},
		},
		{
			name: "prepend within cap",
			cap:  4,
			len:  1,
			got:  make([]any, 0, 4),
			want: []any{1},
		},
		{
			name: "prepend maxing cap",
			cap:  7,
			len:  4,
			got:  []any{2, 3, 4},
			want: []any{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayList := arraylist.Prepend(tt.got, 1)

			if len(arrayList) != tt.len {
				t.Errorf("len is different: got %d - want %d", len(arrayList), tt.len)
			}

			if cap(arrayList) != tt.cap {
				t.Errorf("cap is different: got %d - want %d", cap(arrayList), tt.cap)
			}

			if !reflect.DeepEqual(arrayList, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.got, tt.got, tt.want)
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name string
		cap  int
		len  int
		got  []any
		want []any
	}{
		{
			name: "Insert within cap",
			cap:  7,
			len:  4,
			got:  []any{1, 3, 4},
			want: []any{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayList := arraylist.InsertAt(tt.got, 2, 1)

			if len(arrayList) != tt.len {
				t.Errorf("len is different: got %d - want %d", len(arrayList), tt.len)
			}

			if cap(arrayList) != tt.cap {
				t.Errorf("cap is different: got %d - want %d", cap(arrayList), tt.cap)
			}

			if !reflect.DeepEqual(arrayList, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.got, tt.got, tt.want)
			}
		})
	}
}
