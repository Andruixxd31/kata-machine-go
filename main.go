package main

import (
	"fmt"

	"github.com/andruixxd31/kata-machine-go/src"
)

func main() {
	array := []int{0, 2, 3, 6, 8, 19, 23}
	index := src.BinarySearch(array, 5, 0, len(array))
	fmt.Printf("Index is %d \n", index)
}
