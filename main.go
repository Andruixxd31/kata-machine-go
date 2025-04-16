package main

import (
	"fmt"

	"github.com/andruixxd31/kata-machine-go/src"
)

func main() {
	array := []int{0, 2, 3, 6, 8, 19, 23}
	index := src.BinarySearch(array, 5, 0, len(array))
	index2 := src.BinarySearch(array, 0, 0, len(array))
	index3 := src.BinarySearch(array, 23, 0, len(array))
	index4 := src.BinarySearch(array, 6, 0, len(array))
	fmt.Printf("Index is %d \n", index)
	fmt.Printf("Index is %d \n", index2)
	fmt.Printf("Index is %d \n", index3)
	fmt.Printf("Index is %d \n", index4)
}
