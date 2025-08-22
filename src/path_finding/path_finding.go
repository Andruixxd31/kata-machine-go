package pathfinding

import (
	"reflect"
)

type Point struct {
	X int
	Y int
}

func Solve(maze [][]string, wall string, start Point, end Point) []Point {
	var path []Point
	seen := map[Point]bool{}

	path = Walk(maze, "#", start, end, path, seen)
	return path
}

func Walk(maze [][]string, wall string, position Point, end Point, path []Point, seen map[Point]bool) []Point {
	if position.X < 0 || position.X >= len(maze[0]) || position.Y < 0 || position.Y >= len(maze) {
		return nil
	}

	if maze[position.X][position.Y] == wall {
		return nil
	}

	if _, ok := seen[position]; ok {
		return nil
	}

	if reflect.DeepEqual(position, end) {
		path = append(path, position)
		return path
	}

	seen[position] = true
	path = append(path, position)

	// Walk()
	choices := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for _, choice := range choices {
		result := Walk(maze, "#", Point{position.X + choice[0], position.Y + choice[1]}, end, path, seen)
		if result != nil {
			return result
		}
	}

	path = path[:len(path)-1]

	return nil
}
