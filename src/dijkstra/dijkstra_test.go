package dijkstra_test

import (
	"reflect"
	"testing"

	"github.com/andruixxd31/kata-machine-go/src/dijkstra"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name     string
		graph    dijkstra.Graph
		start    int
		expected map[int]int
	}{
		{
			name: "simple graph",
			graph: dijkstra.Graph{
				0: {{To: 1, Weight: 4}, {To: 2, Weight: 1}},
				1: {{To: 3, Weight: 1}},
				2: {{To: 1, Weight: 2}, {To: 3, Weight: 5}},
				3: {},
			},
			start: 0,
			expected: map[int]int{
				0: 0,
				1: 3,
				2: 1,
				3: 4,
			},
		},
		{
			name: "disconnected graph",
			graph: dijkstra.Graph{
				0: {{To: 1, Weight: 2}},
				1: {},
				2: {{To: 3, Weight: 1}},
				3: {},
			},
			start: 0,
			expected: map[int]int{
				0: 0,
				1: 2,
				2: 1<<63 - 1, // math.MaxInt64
				3: 1<<63 - 1,
			},
		},
		{
			name: "single node graph",
			graph: dijkstra.Graph{
				0: {},
			},
			start: 0,
			expected: map[int]int{
				0: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (heap)", func(t *testing.T) {
			got := dijkstra.Dijkstra(tt.graph, tt.start)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Dijkstra() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
