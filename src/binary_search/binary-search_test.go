package src

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{0, 2, 3, 6, 8, 19, 23}
	t.Run("finds index", func(t *testing.T) {
		got := BinarySearch(array, 2, 0, len(array))
		want := 1

		assertCorrectMessage(t, got, want)
	})

	t.Run("index does not exist", func(t *testing.T) {
		got := BinarySearch(array, 20, 0, len(array))
		want := -1

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
