package main

import (
	"sort"
)

// m.getSolution() get the optimal flow base on number of links in start/end
// and searchSolution based on the flow.
func (m *maze) getSolution() {
	flow := len(m.rooms[m.start].linkTo)
	endLinks := len(m.rooms[m.end].linkTo)
	if endLinks < flow {
		flow = endLinks
	}
	for m.sol == nil && flow > 0 {
		m.searchSolution(flow, 0, []int{})
		flow--
	}
	for _, id := range m.sol.pathIDs {
		m.paths[id].seq = append(m.paths[id].seq, m.end)
	}
	m.sortSolution()
}

// m.searchSolution() goes through all possible combination of paths
// to find the optimal solution (maximum flow, minimum length)
func (m *maze) searchSolution(flow, curID int, pathIDs []int) {
	for ; curID < len(m.paths); curID++ {
		if m.isPathsClash(curID, pathIDs) {
			continue
		}
		newPathID := append([]int{}, pathIDs...)
		newPathID = append(newPathID, curID)
		if len(newPathID) == flow {
			length := 0
			for _, path := range newPathID {
				length += m.paths[path].length
			}
			if m.sol == nil || m.sol.length > length {
				m.sol = &solution{pathIDs: newPathID, length: length}
			}
		} else {
			m.searchSolution(flow, curID+1, newPathID)
		}
	}
}

// m.isPathsClash() check if current path clashes with any paths
// in the current potential solution
func (m *maze) isPathsClash(curID int, pathID []int) bool {
	for _, id := range pathID {
		for _, room := range m.paths[id].seq {
			for _, clash := range m.paths[curID].seq {
				if room == clash {
					return true
				}
			}
		}
	}
	return false
}

// m.sortSolution() sorts m.sol.pathIDs in
// a decreasing order based on path's length
func (m *maze) sortSolution() {
	sort.Slice(m.sol.pathIDs, func(i, j int) bool {
		return m.paths[m.sol.pathIDs[i]].length > m.paths[m.sol.pathIDs[j]].length
	})
}
