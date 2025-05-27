package twocrystallballs_test

import (
	"testing"

	twocrystallballs "github.com/andruixxd31/kata-machine-go/src/two_crystall_balls"
)

func TestTwoCrystalBalls(t *testing.T) {
	tests := []struct {
		name string
		arr  []bool
		want int
	}{
		{
			name: "Index found from a jump",
			arr:  []bool{false, false, true, true},
			want: 2,
		},
		{
			name: "Index found from a jump two",
			arr:  []bool{false, false, false, true, true, true},
			want: 3,
		},
		{
			name: "Index found outside of jump",
			arr:  []bool{false, false, false, false, false, false, false, false, true},
			want: 8,
		},
		{
			name: "Index found outside of jump",
			arr:  []bool{false, false, false, false, false, false, false, false, true, true},
			want: 8,
		},
		{
			name: "No breaking",
			arr:  []bool{false, false, false, false, false, false, false, false, false},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twocrystallballs.TwoCrystralBalls(tt.arr)
			if tt.want != actual {
				t.Errorf("got: %v, want: %v", actual, tt.want)
			}
		})
	}

}
