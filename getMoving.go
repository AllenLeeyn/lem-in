package main

import "fmt"

// m.getMoving() moves ants along each path
func (m *maze) getMoving() {
	nextAntID, prevAnt := 1, 0

	for exitQty, turns := 0, 0; exitQty < m.antQty; turns++ {
		m.movement += fmt.Sprintf("Turn %03d:", turns+1)
		for _, pID := range m.sol.pathIDs {
			if m.paths[pID].antsProcessed == m.paths[pID].antsSet {
				continue
			}
			for i, room := range m.paths[pID].seq {
				if i == 0 && m.paths[pID].antProcessing < m.paths[pID].antsSet {
					m.rooms[room].antID, prevAnt = nextAntID, m.rooms[room].antID
					m.paths[pID].antProcessing++
					nextAntID++
				} else if i == 0 {
					m.rooms[room].antID, prevAnt = 0, m.rooms[room].antID
				} else {
					m.rooms[room].antID, prevAnt = prevAnt, m.rooms[room].antID
				}
				exitQty = m.getMove(room, m.rooms[room].antID, pID, exitQty)
			}
		}
		m.movement += "\n"
	}
}

// m.getMove() records a move/exit if aID is not 0.
func (m *maze) getMove(room string, aID, pID, exitQty int) int {
	if aID != 0 {
		color := colors[aID%12]
		if room == m.end {
			color = "\033[7m" + color
			m.paths[pID].antsProcessed++
			exitQty++
		}
		m.movement += fmt.Sprintf(" %sL%v-%s%s", color, aID, room, resetColor)
	}
	return exitQty
}

// m.getAntAssignment() assigns ants to each path base on path's length
func (m *maze) setAntsToPaths() {

	unSetAnts := m.antQty
	fmt.Println("Unset ants:", unSetAnts)

	numberOfPaths := len(m.sol.pathIDs)

	for unSetAnts > 0 {

		for i, id := range m.sol.pathIDs { // loop thru all optimal paths

			// Calculate current path steps considering length and ants already set
			currentPathSteps := m.paths[id].length + m.paths[id].antsSet
			var nextPathSteps int

			// Check next path if it's not the last one
			if i < numberOfPaths-1 {
				nextID := m.sol.pathIDs[i+1]
				nextPathSteps = m.paths[nextID].length + m.paths[nextID].antsSet
	
			// If the current path steps are less than the next path, assign an ant to the current path
			if currentPathSteps < nextPathSteps {
				m.paths[id].antsSet += 1
				unSetAnts -= 1
				break
			}
		}else { // if the last path is reached without breaking the loop, assign the ant there
			m.paths[id].antsSet += 1
				unSetAnts -= 1
		}
			// Break early if all ants are assigned
			if unSetAnts == 0 {
				break
			}
		}
	}

	// Output final path assignments for debugging
	for _, id := range m.sol.pathIDs {
		fmt.Println(m.paths[id])
	}
}
