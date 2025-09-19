package connectedcomponents_test

import (
	"reflect"
	"testing"

	connectedcomponents "github.com/andruixxd31/kata-machine-go/src/connected_components"
)

// TODO:
// 1.
func TestConnectedComponents(t *testing.T) {
	want := 2

	wantComponents := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 1,
		5: 1,
	}

	graph := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1, 3},
		3: {2},
		4: {5},
		5: {4},
	}

	t.Run("return 3", func(t *testing.T) {
		gotCount, gotComponents := connectedcomponents.CountConnectedComponents(graph)

		assertCount(t, gotCount, want)

		assertComponents(t, gotComponents, wantComponents)
	})
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
