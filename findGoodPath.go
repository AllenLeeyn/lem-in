package main

// Breadth-first search (BFS) to find shortest path; and find an augmenting path
func (l *Lemz) BfsDirected(capacity map[string]map[string]int, flow map[string]map[string]int, parent *map[string]string, usedEndEdges *map[string]bool) bool {
	visited := make(map[string]bool)
	queue := []string{l.start}
	visited[l.start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range l.graph[node] {
			// Check residual capacity in the forward direction
			if !visited[neighbor] && capacity[node][neighbor]-flow[node][neighbor] > 0 {
				// Avoid reusing edges unnecessarily
				if neighbor == l.end && (*usedEndEdges)[node] {
					continue
				}
				(*parent)[neighbor] = node
				visited[neighbor] = true

				if neighbor == l.end {
					(*usedEndEdges)[node] = true
					return true
				}
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

// Edmonds-Karp algorithm to find all paths; all ends must have a path
func (l *Lemz) EdmondsKarp() {
	// Initialize capacity, flow, and usedEndEdges
	capacity := make(map[string]map[string]int)
	flow := make(map[string]map[string]int)
	usedEndEdges := make(map[string]bool)

	for node := range l.graph {
		capacity[node] = make(map[string]int)
		flow[node] = make(map[string]int)
	}

	// Set capacities for all edges
	for _, link := range l.links {
		// Directed graph: set capacity in one direction
		capacity[link[0]][link[1]] = 1
		if _, exists := flow[link[0]]; !exists {
			flow[link[0]] = make(map[string]int)
		}
		if _, exists := flow[link[1]]; !exists {
			flow[link[1]] = make(map[string]int)
		}
	}

	// Use pointers to parent map and usedEndEdges
	for {
		parent := make(map[string]string)
		if !l.BfsDirected(capacity, flow, &parent, &usedEndEdges) {
			break
		}

		// Update flow along the found path
		node := l.end
		for node != l.start {
			prevNode := parent[node]
			flow[prevNode][node]++
			flow[node][prevNode]-- // Residual capacity
			node = prevNode
		}

		// Construct and store the path
		path := []string{}
		node = l.end
		for node != l.start {
			path = append([]string{node}, path...)
			node = parent[node]
		}
		path = append([]string{l.start}, path...)
		l.paths = append(l.paths, path)
	}
}
