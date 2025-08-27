package quicksort_test

import (
	"reflect"
	"testing"

	quicksort "github.com/andruixxd31/kata-machine-go/src/quick_sort"
)

func TestQuickSort(t *testing.T) {
	t.Run("Quick sort test", func(t *testing.T) {
		got := []int{3, 4, 2, 7, 9, 1, 5, 6}
		want := []int{1, 2, 3, 4, 5, 6, 7, 9}
		quicksort.Quicksort(got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v\n", got, want)
		}
	})
}
