package twocrystallballs

import (
	"math"
)

func TwoCrystralBalls(arr []bool) int {
	idx := -1

	jmp := math.Floor(math.Sqrt(float64(len(arr))))

	i := 0
	for i = 0; i < len(arr); i += int(jmp) {
		if arr[i] {
			break
		}
	}

	i -= int(jmp)

	for j := 0; j <= int(jmp) && i+j < len(arr); j++ {
		if arr[i+j] {
			return i + j
		}
	}

	return idx
}
