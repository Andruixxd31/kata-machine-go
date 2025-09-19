package connectedcomponents_test

import (
	"reflect"
	"testing"

	connectedcomponents "github.com/andruixxd31/kata-machine-go/src/connected_components"
)

func TestConnectedComponents(t *testing.T) {
	tests := []struct {
		name           string
		graph          map[int][]int
		wantCount      int
		wantComponents map[int]int
	}{
		{
			name:      "two components",
			wantCount: 2,
			wantComponents: map[int]int{
				0: 0, 1: 0, 2: 0, 3: 0,
				4: 1, 5: 1,
			},
			graph: map[int][]int{
				0: {1},
				1: {0, 2},
				2: {1, 3},
				3: {2},
				4: {5},
				5: {4},
			},
		},
		{
			name:      "single component",
			wantCount: 1,
			wantComponents: map[int]int{
				0: 0, 1: 0, 2: 0, 3: 0,
			},
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1, 3},
				3: {2},
			},
		},
		{
			name:      "disconnected nodes",
			wantCount: 3,
			wantComponents: map[int]int{
				0: 0, 1: 1, 2: 2,
			},
			graph: map[int][]int{
				0: {},
				1: {},
				2: {},
			},
		},
		{
			name:      "multiple small components",
			wantCount: 3,
			wantComponents: map[int]int{
				0: 0, 1: 0, 2: 1, 3: 2,
			},
			graph: map[int][]int{
				0: {1},
				1: {0},
				2: {},
				3: {},
			},
		},
		{
			name:           "empty graph",
			wantCount:      0,
			wantComponents: map[int]int{},
			graph:          map[int][]int{},
		},
		{
			name:      "larger graph",
			wantCount: 3,
			wantComponents: map[int]int{
				0: 0, 1: 0, 2: 0,
				3: 1, 4: 1,
				5: 2,
			},
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
				3: {4},
				4: {3},
				5: {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotComponents := connectedcomponents.CountConnectedComponents(tt.graph)
			assertCount(t, gotCount, tt.wantCount)
			assertComponents(t, gotComponents, tt.wantComponents)
		})
	}
}

func assertCount(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertComponents(t testing.TB, got, want map[int]int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
