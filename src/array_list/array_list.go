package arraylist

// prepend(item: T): void
// insertAt(item: T, idx: number): void
// append(item: T): void
// remove(item: T): T |
// get(idx: number): T | undefined
// removeAt(idx: number): T | undefined

func prepend(arrayList []int, item int) []int {

	return arrayList
}

func Append(arrayList []int, item int) []int {
	n := len(arrayList)
	if n == cap(arrayList) {
		arrayList = increaseCap(arrayList, n)
	}
	arrayList = arrayList[:n+1]
	arrayList[n] = item
	return arrayList
}

func insertAt(arrayList []int, item int, idx int) []int {

	return arrayList
}

func remove(arrayList *[]int) {

}

func get(idx int) {

}

func increaseCap(arrayList []int, n int) []int {
	slice := make([]int, n, n*2+1)
	copy(slice, arrayList)
	arrayList = slice

	return arrayList
}
