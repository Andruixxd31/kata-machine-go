package pathfinding_test

import (
	"fmt"
	"reflect"
	"testing"

	pathfinding "github.com/andruixxd31/kata-machine-go/src/path_finding"
)

func TestPathFinding(t *testing.T) {
	tests := []struct {
		name  string
		maze  [][]string
		start pathfinding.Point
		end   pathfinding.Point
		wall  string
		want  []pathfinding.Point
	}{
		{
			name: "Maze with solution",
			maze: [][]string{
				{"#", "#", "#", "#", "#", "#", "#"},
				{"#", " ", " ", " ", " ", " ", "#"},
				{"#", "#", "#", " ", "#", " ", "#"},
				{"#", " ", " ", " ", "#", " ", "#"},
				{"#", "#", "#", "#", "#", " ", "#"},
				{"#", " ", " ", " ", " ", " ", "#"},
				{"#", "#", "#", "#", "#", "#", "#"},
			},
			want:  []pathfinding.Point{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {5, 4}, {5, 3}, {5, 2}, {5, 1}},
			start: pathfinding.Point{X: 1, Y: 1},
			end:   pathfinding.Point{X: 5, Y: 1},
			wall:  "#",
		},
		{
			name: "Maze with no solution",
			maze: [][]string{
				{"#", "#", "#", "#", "#", "#", "#"},
				{"#", " ", " ", " ", " ", " ", "#"},
				{"#", "#", "#", " ", "#", " ", "#"},
				{"#", " ", " ", " ", "#", " ", "#"},
				{"#", "#", "#", "#", "#", " ", "#"},
				{"#", " ", "#", " ", " ", " ", "#"},
				{"#", "#", "#", "#", "#", "#", "#"},
			},
			want:  []pathfinding.Point{},
			start: pathfinding.Point{X: 1, Y: 1},
			end:   pathfinding.Point{X: 5, Y: 1},
			wall:  "#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathfinding.Solve(tt.maze, tt.wall, tt.start, tt.end)

			if got == nil {
				got = []pathfinding.Point{}
			}

			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println(reflect.TypeOf(got))
				fmt.Println(reflect.TypeOf(tt.want))
				t.Errorf("Solve() = %+v, want %+v", got, tt.want)
			}
		})
	}

}
