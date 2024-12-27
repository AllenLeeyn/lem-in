package main

// Helper function to perform BFS and find an augmenting path
func (l *Lemz) Bfs(capacity map[string]map[string]int, flow map[string]map[string]int, parent map[string]string) bool {
	visited := make(map[string]bool)
	queue := []string{l.start}
	visited[l.start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range l.graph[node] {
			if !visited[neighbor] && capacity[node][neighbor]-flow[node][neighbor] > 0 { // There's remaining capacity
				parent[neighbor] = node
				if neighbor == l.end {
					return true
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}

// Edmonds-Karp algorithm to find all paths
func (l *Lemz) EdmondsKarp() {
	// Initialize flow and capacity maps
	capacity := make(map[string]map[string]int)
	flow := make(map[string]map[string]int)

	for node := range l.graph {
		capacity[node] = make(map[string]int)
		flow[node] = make(map[string]int)
	}

	// Set capacity for each edge
	for _, link := range l.links {
		capacity[link[0]][link[1]] = 1
		capacity[link[1]][link[0]] = 1 // For undirected graph
	}

	// Find augmenting paths using BFS
	for {
		parent := make(map[string]string)
		if !l.Bfs(capacity, flow, parent) {
			break // No more augmenting paths
		}

		// Find the bottleneck capacity (minimum residual capacity along the path)
		pathFlow := 1 // For unweighted graphs, flow is 1 for each path
		node := l.end
		for node != l.start {
			prevNode := parent[node]
			flow[prevNode][node] += pathFlow
			flow[node][prevNode] -= pathFlow
			node = prevNode
		}

		// Add the found path to paths
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

/* DOESNT WORK: // list of path single values in order from start to end [but this is undirected flow!!!!]
func (l *Lemz) GetPathSingleOrder(links [][2]string) [][]string {
	pathSingleOrder := []string{}
	editedLinks := l.links
	from := l.start
	fmt.Printf("starting l.links: %#v\n", l.links)

	for i := 0; i < l.EndCount(l.links); i++ {
		fmt.Printf("endcount: %v\n", l.EndCount(l.links))
		for i := 0; i < len(editedLinks); i++ {

			fmt.Printf("FOR:editedLinks[%v]: %v\n", i, editedLinks[i])

			if editedLinks[i][0] == from {
				pathSingleOrder = append(pathSingleOrder, editedLinks[i][1])
				from = editedLinks[i][1]
				editedLinks = append(editedLinks[:i], editedLinks[i+1:]...)
				continue
			} else if l.links[i][1] == from {
				pathSingleOrder = append(pathSingleOrder, l.links[i][0])
				from = l.links[i][0]
				l.links = append(l.links[:i], l.links[i+1:]...)
				fmt.Printf("pathSingleOrder: %#v | from: %#v\n", pathSingleOrder, from)
				continue
			}

			if from == l.end {
				fmt.Printf("from == l.end\n")
				editedLinks = l.links
				from = l.start
				l.pathsSingleOrder = append(l.pathsSingleOrder, pathSingleOrder)
				break
			}

			fmt.Printf("FOR: pathSingleOrder: %#v | from: %#v\n", pathSingleOrder, from)

		}
	}
	fmt.Printf("PathSSSingleOrder: %#v || len: %#v\n", l.pathsSingleOrder, len(l.pathsSingleOrder))
	return l.pathsSingleOrder
}

func (l *Lemz) EndCount(links [][2]string) int {
	endCount := 0
	for _, link := range l.links {
		if link[0] == l.end || link[1] == l.end {
			endCount++
		}
	}
	return endCount
}
*/
