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

func (m *maze) getOptimize(flow, cur int, pathID []int) {
	for i := cur; i < len(m.paths); i++ {
		if m.isPathsClash(i, pathID) {
			continue
		}
		newPathID := append(pathID, i)

		if len(newPathID) == flow {
			length := 0
			for _, path := range newPathID {
				length += m.paths[path].length
			}
			if m.sol == nil || m.sol.length > length {
				m.sol = &solution{pathID: newPathID, length: length}
			}
		} else {
			m.getOptimize(flow, i+1, newPathID)
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
	for i := 0; i < len(m.sol.pathID)-1; i++ {
		cur, next := m.sol.pathID[i], m.sol.pathID[i+1]
		if m.paths[cur].length < m.paths[next].length {
			m.sol.pathID[i], m.sol.pathID[i+1] =
				m.sol.pathID[i+1], m.sol.pathID[i]
		}
	}
}
