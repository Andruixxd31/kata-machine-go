package dijkstra

import (
	"container/heap"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph adjacency list
type Graph map[int][]Edge

type Node struct {
	Value    int
	Distance int
	Index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Distance < pq[j].Distance }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].Index = i; pq[j].Index = j }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
func (pq *PriorityQueue) update(item *Node, distance int) {
	item.Distance = distance
	heap.Fix(pq, item.Index)
}

func Dijkstra(graph Graph, start int) map[int]int {
	//TODO:
	// init the pq
	// Set all distances to math.MaxValue
	// push start node
	// iterate
	// 		pop node from pq
	// 		from node get all neighbours and update distances if less than before and call heap update
	//
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt64
	}

	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{Value: start, Distance: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node)

		u := cur.Value
		d := cur.Distance

		if d > dist[u] {
			continue
		}

		for _, edge := range graph[u] {
			v := edge.To
			newDist := dist[u] + edge.Weight
			if newDist < dist[v] {
				dist[v] = newDist
				heap.Push(&pq, &Node{Value: v, Distance: newDist})
			}

		}

	}

	return dist
}
