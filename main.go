package main

import (
	"fmt"

	twocrystallballs "github.com/andruixxd31/kata-machine-go/src/two_crystall_balls"
)

func main() {
	// array := []bool{false, false, false, false, false, true, true, true, true}
	array := []bool{false, false, true, true, true, true}
	// array := []bool{false, false, false, false, false, true, true, true, true}
	index := twocrystallballs.TwoCrystralBalls(array)
	fmt.Printf("Index is %d \n", index)
}
