package src

// if the index is found returns the index in the arr else returns -1
// Implemented using recurrence
// Will implement a search of O(logn)
func BinarySearch(arr []int, target, low, high int) int {
	pivot := low + (high-low)/2

	if low > high {
		return -1
	}

	if arr[pivot] > target {
		return BinarySearch(arr, target, low, pivot-1)
	} else if arr[pivot] < target {
		return BinarySearch(arr, target, pivot+1, high)
	} else {
		return pivot
	}
}
