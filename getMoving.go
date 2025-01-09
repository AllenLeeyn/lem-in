package main

import "fmt"

// m.getMoving() moves ants along each path
func (m *maze) getMoving() {
	nextAntID, prevAnt := 1, 0
	antsExited, turns := 0, 0

	for ; antsExited < m.antQty; turns++ {
		m.movement += fmt.Sprintf("Turn %03d:", turns+1)
		for _, pathID := range m.optimalPaths {
			if m.paths[pathID].antsProcessed == m.paths[pathID].antsSet {
				continue
			}
			for i, room := range m.paths[pathID].path {
				if i == 0 && m.paths[pathID].antProcessing < m.paths[pathID].antsSet { //if it is the first room, and there is still ants left to process
					m.rooms[room].antID, prevAnt = nextAntID, m.rooms[room].antID //
					m.paths[pathID].antProcessing++
					nextAntID++
				} else if i == 0 {
					m.rooms[room].antID, prevAnt = 0, m.rooms[room].antID
				} else {
					m.rooms[room].antID, prevAnt = prevAnt, m.rooms[room].antID
				}
				antsExited = m.getMove(room, m.rooms[room].antID, pathID, antsExited)
			}
		}
		m.movement += "\n"
	}
	//m.movement += fmt.Sprintf("\nNumber of ants exited: %d", antsExited) + "\n"
}

// m.getMove() records a move/exit if aID is not 0.
func (m *maze) getMove(room string, antID, pathID, antsExited int) int {
	if antID != 0 {
		color := colors[antID%12]
		if room == m.end {
			color = "\033[7m" + color
			m.paths[pathID].antsProcessed++
			antsExited++
		}
		m.movement += fmt.Sprintf(" %sL%d-%s%s", color, antID, room, resetColor)
	}
	return antsExited
}

// m.getAntAssignment() assigns ants to each path base on path's length
func (m *maze) setAntsToPaths() {

	unSetAnts := m.antQty
	numberOfPaths := len(m.optimalPaths)

	for unSetAnts > 0 {
		for i, id := range m.optimalPaths { // loop thru all optimal paths

			// Calculate current path steps considering length and ants already set
			currentPathSteps := m.paths[id].length + m.paths[id].antsSet
			var nextPathSteps int

			// Check next path if it's not the last one
			if i < numberOfPaths-1 {
				nextID := m.optimalPaths[i+1]
				nextPathSteps = m.paths[nextID].length + m.paths[nextID].antsSet

				// If the current path steps are less than the next path, assign an ant to the current path
				if currentPathSteps < nextPathSteps {
					m.paths[id].antsSet += 1
					unSetAnts -= 1
					break
				}
			} else { // if the last path is reached without breaking the loop, assign the ant there
				m.paths[id].antsSet += 1
				unSetAnts -= 1
			}
			// Break early if all ants are assigned
			if unSetAnts == 0 {
				break
			}
		}
	}
}
