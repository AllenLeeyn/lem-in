package main

func (m *maze) getSolution() {
	flow := len(m.rooms[m.start].linkTo)
	endLinks := len(m.rooms[m.end].linkTo)
	if endLinks < flow {
		flow = endLinks
	}
	for m.sol == nil && flow > 0 {
		m.getOptimize(flow, 0, []int{})
		flow--
	}
	m.getSolutionSorted()
}

func (m *maze) getOptimize(flow, cur int, paths []int) {
	for i := cur; i < len(m.paths); i++ {
		if m.isPathsClash(i, paths) {
			continue
		}
		newPaths := append(paths, i)

		if len(newPaths) == flow {
			length := 0
			for _, path := range newPaths {
				length += m.paths[path].length
			}
			if m.sol == nil || m.sol.length > length {
				m.sol = &solution{paths: newPaths, length: length}
			}
		} else {
			m.getOptimize(flow, i+1, newPaths)
		}
	}
}

func (m *maze) isPathsClash(cur int, sol []int) bool {
	for _, path := range sol {
		for _, room := range m.paths[path].seq {
			for _, clash := range m.paths[cur].seq {
				if room == clash {
					return true
				}
			}
		}
	}
	return false
}

func (m *maze) getSolutionSorted() {
	for i := 0; i < len(m.sol.paths)-1; i++ {
		cur, next := m.sol.paths[i], m.sol.paths[i+1]
		if m.paths[cur].length < m.paths[next].length {
			m.sol.paths[i], m.sol.paths[i+1] =
				m.sol.paths[i+1], m.sol.paths[i]
		}
	}
}
