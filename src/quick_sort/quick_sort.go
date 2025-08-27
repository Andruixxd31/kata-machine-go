package quicksort

func Quicksort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[hi], arr[idx] = arr[idx], arr[hi]

	return idx
}

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(arr, lo, hi)
	qs(arr, lo, pivot-1)
	qs(arr, pivot+1, hi)

}
