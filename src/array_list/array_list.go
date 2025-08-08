package arraylist

func Prepend(arrayList []int, item int) []int {
	n := len(arrayList)
	if n == cap(arrayList) {
		arrayList = increaseCap(arrayList, n)
	}

	arrayList = arrayList[:n+1]
	copy(arrayList[1:], arrayList[:])
	arrayList[0] = item
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

func InsertAt(arrayList []int, item int, idx int) []int {
	n := len(arrayList)
	if n == cap(arrayList) {
		arrayList = increaseCap(arrayList, n)
	}

	arrayList = arrayList[:n+1]
	copy(arrayList[idx+1:], arrayList[idx:])
	arrayList[idx] = item
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
