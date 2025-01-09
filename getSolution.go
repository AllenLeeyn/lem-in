package main

import (
	"math"
)

// m.getSolution() get the optimal flow based on number of links in start/end
// and searchSolution based on the flow.
func (m *maze) getSolution() {
	maxFlow := len(m.rooms[m.start].linkTo)
	endLinks := len(m.rooms[m.end].linkTo)
	if endLinks < maxFlow {
		maxFlow = endLinks
	}
	m.optimalPaths = []int{0}
	m.searchSolution(maxFlow, 0, []int{})

	for _, id := range m.optimalPaths {
		m.paths[id].path = append(m.paths[id].path, m.end)
	}
}

// m.searchSolution() search for combinations of paths based on flow.
func (m *maze) searchSolution(flow, curID int, pathIDs []int) {
	for ; curID < len(m.paths); curID++ {
		lengthLimit := m.getLengthLimit(pathIDs, flow)
		if (m.paths[curID].length - 1) > lengthLimit {
			return
		}
		if m.isPathsClash(curID, pathIDs) {
			continue
		}
		newPathIDs := append([]int{}, pathIDs...)
		newPathIDs = append(newPathIDs, curID)
		length := 0
		for _, path := range newPathIDs {
			length += m.paths[path].length
		}
		estNumSteps := m.getStepEstimate(newPathIDs)
		if estNumSteps != 0 && estNumSteps < m.getStepEstimate(m.optimalPaths) { //if newPathIDs estimate steps is smaller than the current best solution
			m.optimalPaths = newPathIDs // assign newPathIDs to the current best
		}
		m.searchSolution(flow, curID+1, newPathIDs)
	}
}

// m.getLengthLimit() calculates the average length per path
// based on the number of rooms and the number of paths.
// This is an estimate length to limit search.
func (m *maze) getLengthLimit(pathIDs []int, flow int) int {
	numberOfRooms := len(m.rooms) // number of all rooms
	length := 0
	for _, id := range pathIDs {
		length += m.paths[id].length - 1 // number of rooms in current solution
	}
	remainingLength := float64(numberOfRooms - length)
	remainingPath := float64(flow - len(pathIDs))          // number of paths we are looking - number of paths collected already
	return int(math.Ceil(remainingLength / remainingPath)) // we can eliminate all paths that are longer than this
}

// m.stepEstimate() calculates an estimate number of turns.
// This help us to compare slutions.
func (m *maze) getStepEstimate(pathIDs []int) int {
	pathsCnt := len(pathIDs)
	longestLength := m.paths[pathIDs[pathsCnt-1]].length

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
func (m *maze) isPathsClash(curID int, pathIDs []int) bool {
	for _, id := range pathIDs {
		for _, room := range m.paths[id].path {
			for _, clash := range m.paths[curID].path {
				if room == clash {
					return true
				}
			}
		}
	}
	return false
}
