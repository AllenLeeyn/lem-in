package main

import (
	"math"
)

// m.getSolution() get the optimal flow base on number of links in start/end
// and searchSolution based on the flow.
func (m *maze) getSolution() {
	maxFlow := len(m.rooms[m.start].linkTo)
	endLinks := len(m.rooms[m.end].linkTo)
	if endLinks < maxFlow {
		maxFlow = endLinks
	}
	m.sol = &solution{pathIDs: []int{0}, length: len(m.paths[0].seq)}
	for curFlow := 1; curFlow <= maxFlow; curFlow++ {
		m.searchSolution(curFlow, 0, []int{})
	}
	for _, id := range m.sol.pathIDs {
		m.paths[id].seq = append(m.paths[id].seq, m.end)
	}
}

// m.searchSolution() search for combinations of paths based on flow.
func (m *maze) searchSolution(flow, curID int, pathIDs []int) {
	for ; curID < len(m.paths); curID++ {
		lengthLimit := m.getLengthLimit(pathIDs, flow)
		if m.paths[curID].length-1 > lengthLimit {
			return
		}
		if m.isPathsClash(curID, pathIDs) {
			continue
		}
		newPathIDs := append([]int{}, pathIDs...)
		newPathIDs = append(newPathIDs, curID)
		if len(newPathIDs) == flow {
			length := 0
			for _, path := range newPathIDs {
				length += m.paths[path].length
			}
			estNumSteps := m.evalSolution(newPathIDs)
			if estNumSteps != 0 && m.evalSolution(m.sol.pathIDs) > estNumSteps {
				m.sol = &solution{pathIDs: newPathIDs, length: length}
			}
			return
		} else {
			m.searchSolution(flow, curID+1, newPathIDs)
		}
	}
}

// m.getLengthLimit() calculates the average length per path
// based on the number of rooms and the number of paths.
// This is an estimate length to limit search.
func (m *maze) getLengthLimit(pathIDs []int, flow int) int {
	lenMaze := len(m.rooms)
	length := 0
	for _, id := range pathIDs {
		length += m.paths[id].length - 1
	}
	remainingLength := float64(lenMaze - length)
	remainingPath := float64(flow - len(pathIDs))
	return int(math.Ceil(remainingLength / remainingPath))
}

// m.evalSolution() calculates an estimate number of turns.
// This help us to compare slutions.
func (m *maze) evalSolution(pathIDs []int) int {
	pathsCnt := len(pathIDs)
	longestPath := pathIDs[pathsCnt-1]
	longestLength := m.paths[longestPath].length

	diff := 0
	for _, id := range pathIDs {
		diff += longestLength - m.paths[id].length
	}
	if m.antQty-diff <= 0 {
		return 0
	}
	antsPerPath := math.Ceil(float64(m.antQty-diff-pathsCnt) / float64(pathsCnt))
	return int(antsPerPath) + longestLength
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
