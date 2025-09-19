package connectedcomponents

func CountConnectedComponents(graph map[int][]int) (int, map[int]int) {
	// TODO:
	// iterate through each node (graph key)
	// do dfs on key
	// check if visited, if not update count
	// check visited
	// change key to count
	//

	visited := make([]bool, len(graph))
	component := make(map[int]int)
	count := 0

	for k := range len(graph) {
		if visited[k] {
			continue
		}

		dfs(graph, component, k, visited, count)
		count++
	}

	return count, component
}

func dfs(graph map[int][]int, component map[int]int, node int, visited []bool, count int) {
	if visited[node] {
		return
	}
	visited[node] = true
	component[node] = count

	for _, child := range graph[node] {
		dfs(graph, component, child, visited, count)
	}
}
