package arraylist

func Prepend(arrayList []interface{}, item interface{}) []interface{} {
	n := len(arrayList)
	if n == cap(arrayList) {
		arrayList = increaseCap(arrayList, n)
	}

	arrayList = arrayList[:n+1]
	copy(arrayList[1:], arrayList[:])
	arrayList[0] = item
	return arrayList
}

func Append(arrayList []interface{}, item interface{}) []interface{} {
	n := len(arrayList)
	if n == cap(arrayList) {
		arrayList = increaseCap(arrayList, n)
	}

	arrayList = arrayList[:n+1]
	arrayList[n] = item
	return arrayList
}

func InsertAt(arrayList []interface{}, item int, idx int) []interface{} {
	n := len(arrayList)
	if n == cap(arrayList) {
		arrayList = increaseCap(arrayList, n)
	}

	arrayList = arrayList[:n+1]
	copy(arrayList[idx+1:], arrayList[idx:])
	arrayList[idx] = item
	return arrayList
}

func increaseCap(arrayList []interface{}, n int) []interface{} {
	slice := make([]interface{}, n, n*2+1)
	copy(slice, arrayList)
	arrayList = slice

	return arrayList
}
